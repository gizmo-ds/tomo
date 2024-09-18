/*
 * Copyright (c) 2024 Gizmo.
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package app

import (
	"os"
	"path/filepath"

	"github.com/duke-git/lancet/v2/xerror"
)

const DataPath = "./appdata"

var (
	CacheDir = filepath.Join(DataPath, "cache")
)

func init() {
	_ = os.MkdirAll(DataPath, 0750)
	_ = os.MkdirAll(CacheDir, 0740)
	cacheDatabaseInit()
}

func WebviewUserDataPath() string {
	return xerror.TryUnwrap(filepath.Abs(DataPath))
}
