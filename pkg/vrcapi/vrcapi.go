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
	"net/url"

	"github.com/go-resty/resty/v2"
)

var (
	ErrNoCookieJar        = errors.New("no cookie jar")
	ErrAuthCookieNotFound = errors.New("auth cookie not found")
)

type common struct {
	httpClient *resty.Client
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

	AuthAPI
}

func New(opts ...Option) *VRCAPI {
	options := defaultOptions()
	for _, opt := range opts {
		opt.apply(options)
	}
	api := &VRCAPI{
		common:  options,
		AuthAPI: &authAPI{common: options},
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
