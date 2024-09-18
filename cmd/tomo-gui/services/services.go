/*
 * Copyright (c) 2024 Gizmo.
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package services

import "github.com/wailsapp/wails/v3/pkg/application"

var MainWindow application.Window

func Services() []application.Service {
	return []application.Service{
		application.NewService(&AppService{}),
		application.NewService(&FileService{}),
	}
}
