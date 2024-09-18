/*
 * Copyright (c) 2024 Gizmo.
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package services

import (
	"encoding/base64"

	"github.com/gizmo-ds/tomo/cmd/tomo-gui/app"
)

type FileService struct{}

func (FileService) GetThumbHash(pathname string) string {
	hash, err := app.GetThumbHashWithPathname(pathname)
	if err != nil {
		return ""
	}
	return base64.StdEncoding.EncodeToString(hash)
}
