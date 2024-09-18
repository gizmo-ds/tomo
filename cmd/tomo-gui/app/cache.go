/*
 * Copyright (c) 2024 Gizmo.
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package app

import (
	"bytes"
	"errors"
	"fmt"
	"image"
	"image/png"
	"io"
	"log/slog"
	"mime"
	"net/http"
	"os"
	"path/filepath"
	"slices"
	"strings"

	"github.com/duke-git/lancet/v2/slice"
	"github.com/galdor/go-thumbhash"
	"github.com/gizmo-ds/tomo/pkg/vrcapi"
	"go.etcd.io/bbolt"
)

var (
	cacheDatabase *bbolt.DB
)

const (
	vrchatFilePrefix = "/@vrchat-file/"
	vrchatFileServer = "https://api.vrchat.cloud/"

	thumbhashBucket = "thumbhash"
)

func cacheDatabaseInit() {
	var err error
	cacheDatabase, err = bbolt.Open(filepath.Join(DataPath, "cache.db"), 0600, nil)
	if err != nil {
		slog.Error("failed to open cache db", "error", err.Error())
		os.Exit(1)
	}
	_ = cacheDatabase.View(func(tx *bbolt.Tx) error {
		_, err := tx.CreateBucketIfNotExists([]byte(thumbhashBucket))
		return err
	})
	OnShutdown(func() { _ = cacheDatabase.Close() })
}

type filePathInfo struct {
	Type string
	ID   string
	Size string
}

func getFilePathInfo(u string) (*filePathInfo, error) {
	u = strings.TrimPrefix(u, vrchatFilePrefix)
	parts := strings.Split(u, "/")
	if len(parts) < 6 {
		slog.Warn("invalid file path", "path", u)
		return nil, errors.New("invalid file path")
	}
	info := &filePathInfo{
		Type: parts[2],
		ID:   parts[3],
		Size: parts[4],
	}
	if info.Type != "file" && info.Type != "image" {
		slog.Warn("invalid file path", "path", u)
		return nil, errors.New("invalid file path")
	}
	return info, nil
}

func StoreThumbHash(key string, hash []byte) error {
	return cacheDatabase.Update(func(tx *bbolt.Tx) error {
		b, err := tx.CreateBucketIfNotExists([]byte(thumbhashBucket))
		if err != nil {
			return err
		}
		return b.Put([]byte(key), hash)
	})
}

func GetThumbHash(key string) ([]byte, error) {
	var hash []byte
	err := cacheDatabase.View(func(tx *bbolt.Tx) error {
		b := tx.Bucket([]byte(thumbhashBucket))
		if b == nil {
			return errors.New("bucket not found")
		}
		hash = b.Get([]byte(key))
		return nil
	})
	return hash, err
}

func GetThumbHashWithPathname(pathname string) ([]byte, error) {
	info, err := getFilePathInfo(pathname)
	if err != nil {
		return nil, err
	}
	key := strings.Join([]string{info.Type, info.ID, info.Size}, "_")
	hash, err := GetThumbHash(key)
	if err != nil || hash == nil {
		filename := filepath.Join(CacheDir, key) + ".png"
		if _, err = os.Stat(filename); err == nil {
			f, err := os.Open(filename)
			if err != nil {
				return nil, err
			}
			defer f.Close()

			hash, err = encodeThumbHash([]string{".png"}, f)
			if err != nil {
				return nil, err
			}
			if err = StoreThumbHash(key, hash); err != nil {
				return nil, err
			}
		}
	}
	return hash, nil
}

func encodeThumbHash(extensions []string, r io.Reader) ([]byte, error) {
	var err error
	var img image.Image
	if slices.Contains(extensions, ".png") {
		img, err = png.Decode(r)
		if err != nil {
			return nil, err
		}
	} else {
		return nil, fmt.Errorf("unsupported image type: %v", extensions)
	}
	return thumbhash.EncodeImage(img), nil
}

func CacheMiddlewareHandler(next http.Handler) http.Handler {
	api := vrcapi.New()
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.HasPrefix(r.URL.Path, vrchatFilePrefix) {
			info, err := getFilePathInfo(r.URL.Path)
			if err == nil {
				fileKey := strings.Join([]string{info.Type, info.ID, info.Size}, "_")
				filename := filepath.Join(CacheDir, fileKey) + ".png"
				if _, err := os.Stat(filename); errors.Is(err, os.ErrNotExist) {
					downloadUrl := strings.Replace(r.URL.Path, vrchatFilePrefix, vrchatFileServer, 1)

					resp, err := api.HttpDownload(downloadUrl)
					if err != nil {
						slog.Error("failed to download file", "error", err.Error())
						return
					}
					extensions, err := mime.ExtensionsByType(resp.Header().Get("Content-Type"))
					if err == nil {
						extensions = slice.Intersection(extensions, []string{".png"})
					}
					var buf *bytes.Buffer

					var writers []io.Writer
					if len(extensions) > 0 {
						buf = bytes.NewBuffer(nil)
						writers = append(writers, buf)
					}

					f, err := os.Create(filename)
					if err != nil {
						slog.Error("failed to create cache file", "error", err.Error())
						return
					}
					defer f.Close()
					writers = append(writers, f)

					_, err = io.Copy(io.MultiWriter(writers...), resp.RawBody())
					if err != nil {
						slog.Error("failed to write cache file", "error", err.Error())
						return
					}

					hash, err := encodeThumbHash(extensions, buf)
					if err != nil {
						if err = StoreThumbHash(fileKey, hash); err != nil {
							slog.Error("failed to store thumb hash", "error", err.Error())
							return
						}
					}
				}
				http.ServeFile(w, r, filename)
				return
			}
		}
		next.ServeHTTP(w, r)
	})
}
