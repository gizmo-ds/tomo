//go:build linux

/*
 * Copyright (c) 2024 Gizmo.
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package app

import (
	"github.com/jaypipes/ghw/pkg/gpu"
	"github.com/jaypipes/ghw/pkg/option"
	"os"
	"strings"
)

func init() {
	const nvidia = "nvidia" // 🐧🖕️

	if info, _ := gpu.New(option.WithNullAlerter()); info != nil {
		for _, card := range info.GraphicsCards {
			if card.DeviceInfo != nil && card.DeviceInfo.Vendor != nil &&
				(strings.Contains(strings.ToLower(card.DeviceInfo.Vendor.Name), nvidia) ||
					strings.ToLower(card.DeviceInfo.Driver) == nvidia) {
				// https://bugs.webkit.org/show_bug.cgi?id=254807
				_ = os.Setenv("WEBKIT_DISABLE_DMABUF_RENDERER", "1")
			}
		}
	}
}
