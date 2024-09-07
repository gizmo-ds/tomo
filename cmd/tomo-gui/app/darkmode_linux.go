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
	"os/exec"
	"strings"
)

func IsCurrentlyDarkMode() bool {
	path, err := exec.LookPath("gsettings")
	if err == nil {
		out, err := exec.Command(path, "get", "org.gnome.desktop.interface", "gtk-theme").Output()
		if err != nil {
			return false
		}
		theme := strings.TrimRightFunc(strings.ToLower(string(out)), func(r rune) bool { return r == '\'' || r == '\n' })
		return strings.HasSuffix(theme, "dark")
	}
	return false
}
