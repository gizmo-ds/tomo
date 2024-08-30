/*
 * Copyright (c) 2024 Gizmo.
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package main

import (
	"embed"
	"log/slog"
	"os"
	"path/filepath"

	"github.com/duke-git/lancet/v2/condition"
	"github.com/duke-git/lancet/v2/xerror"
	"github.com/gizmo-ds/tomo/cmd/tomo-gui/app"
	"github.com/gizmo-ds/tomo/cmd/tomo-gui/services"
	"github.com/wailsapp/wails/v3/pkg/application"
)

//go:embed assets
var assets embed.FS

func main() {
	a := application.New(application.Options{
		Name:     app.Name,
		Services: services.Services(),
		Assets:   application.AssetOptions{Handler: application.AssetFileServerFS(assets)},
		Windows:  application.WindowsOptions{WebviewUserDataPath: xerror.TryUnwrap(filepath.Abs("./appdata"))},
	})

	a.NewWebviewWindowWithOptions(application.WebviewWindowOptions{
		Title: app.Name,
		BackgroundColour: condition.TernaryOperator(
			app.IsDarkMode(),
			application.NewRGB(0x10, 0x10, 0x14),
			application.NewRGB(0xff, 0xff, 0xff)),
		URL:               "/",
		EnableDragAndDrop: false,
		Centered:          true,
		Width:             900,
		Height:            600,
	})

	if err := a.Run(); err != nil {
		slog.Error("Failed to run app", "error", err)
		os.Exit(1)
	}
}
