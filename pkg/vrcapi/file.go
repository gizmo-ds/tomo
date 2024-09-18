/*
 * Copyright (c) 2024 Gizmo.
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package vrcapi

import (
	"errors"
	"log/slog"
	"net/http"

	"github.com/go-resty/resty/v2"
)

var ErrDownloadFailed = errors.New("download failed")

func (c *common) HttpDownload(u string) (*resty.Response, error) {
	resp, err := c.httpClient.R().SetDoNotParseResponse(true).Get(u)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode() != http.StatusOK {
		slog.Error("download failed", "status", resp.StatusCode(), "body", resp.String())
		return nil, ErrDownloadFailed
	}
	return resp, nil
}
