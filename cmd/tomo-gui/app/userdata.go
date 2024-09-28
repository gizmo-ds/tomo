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
)

const UserDataPath = "./userdata"

func init() {
	_ = os.MkdirAll(UserDataPath, 0700)
}