/*
 * Copyright (c) 2024 Gizmo
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package vrcapi

import (
	"errors"
	"net/http"
	"net/url"

	"github.com/go-resty/resty/v2"
)

var (
	ErrNoCookieJar        = errors.New("no cookie jar")
	ErrAuthCookieNotFound = errors.New("auth cookie not found")
	ErrRequestFailed      = errors.New("request failed")
)

type common struct {
	httpClient *resty.Client
	debugMode  bool
}

func defaultOptions() *common {
	return &common{
		httpClient: resty.New().
			SetHeader("User-Agent", DefaultUserAgent).
			SetBaseURL(DefaultEndpoint),
	}
}

type VRCAPI struct {
	*common

	AuthAPI AuthAPI
	UserAPI UserAPI
}

func New(opts ...Option) *VRCAPI {
	options := defaultOptions()
	for _, opt := range opts {
		opt.apply(options)
	}
	api := &VRCAPI{
		common:  options,
		AuthAPI: &authAPI{common: options},
		UserAPI: &userAPI{common: options},
	}
	return api
}

func (c *common) SetProxy(proxy string) {
	c.httpClient.SetProxy(proxy)
}

func (c *common) AuthCookie() (string, error) {
	jar := c.httpClient.GetClient().Jar
	if jar == nil {
		return "", ErrNoCookieJar
	}
	u, err := url.Parse(c.httpClient.BaseURL)
	if err != nil {
		return "", err
	}
	cookies := jar.Cookies(u)
	for _, cookie := range cookies {
		if cookie.Name == "auth" {
			return cookie.Value, nil
		}
	}
	return "", ErrAuthCookieNotFound
}

func (c *common) SetAuthCookie(cookie string) {
	c.httpClient.SetCookie(&http.Cookie{Name: "auth", Value: cookie})
}

func (c *common) SetDebugMode(debug bool) {
	c.debugMode = debug
	c.httpClient.SetDebug(debug)
}
