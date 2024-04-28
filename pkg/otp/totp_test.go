/*
 * Copyright (c) 2024 Gizmo.
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package otp

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTOTP(t *testing.T) {
	code, err := TOTP("GM4VC2CQN5UGS33ZJJVWYUSFMQ4HOQJW", 1662681600)
	require.NoError(t, err)
	require.Equal(t, uint32(473526), code)
}
