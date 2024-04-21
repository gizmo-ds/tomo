/*
 * Copyright (c) 2024 Gizmo.
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package vrcapi

import (
	"github.com/go-resty/resty/v2"
)

type Option interface {
	apply(*common)
}

type funcOption struct {
	fn func(*common)
}

func (fo *funcOption) apply(do *common) {
	fo.fn(do)
}

func newFuncOption(fn func(*common)) *funcOption {
	return &funcOption{fn: fn}
}

func WithUserAgent(userAgent string) Option {
	return newFuncOption(func(o *common) {
		o.httpClient.SetHeader("User-Agent", userAgent)
	})
}
func WithHTTPClient(httpClient *resty.Client) Option {
	return newFuncOption(func(o *common) { o.httpClient = httpClient })
}
func WithProxy(proxy string) Option {
	return newFuncOption(func(o *common) {
		o.httpClient.SetProxy(proxy)
	})
}
