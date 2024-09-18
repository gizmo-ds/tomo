/*
 * Copyright (c) 2024 Gizmo.
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package vrcapi

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

var api *VRCAPI

func TestMain(m *testing.M) {
	_ = godotenv.Load()
	api = New()
	if proxy := os.Getenv("PROXY_URL"); proxy != "" {
		api.SetProxy(proxy)
	}
	m.Run()
}
