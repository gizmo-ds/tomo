/*
 * Copyright (c) 2024 Gizmo.
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package utils

import "golang.org/x/exp/constraints"

func Clamp[T constraints.Ordered](value, min, max, defaultValue T) T {
	if value < min {
		return defaultValue
	}
	if value > max {
		return max
	}
	return value
}
