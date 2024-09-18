/*
 * Copyright (c) 2024 Gizmo.
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package vrcapi

import (
	"errors"
	"os"
	"slices"
	"strconv"
	"strings"
	"testing"

	"github.com/gizmo-ds/tomo/pkg/otp"
	"github.com/stretchr/testify/require"
)

func TestAuthAPI_Login(t *testing.T) {
	username, password := os.Getenv("VRC_USERNAME"), os.Getenv("VRC_PASSWORD")
	if username == "" || password == "" {
		t.Skip("VRC_USERNAME and VRC_PASSWORD not set")
	}
	currentUser, err := api.AuthAPI.Login(username, password)
	var err2fa RequiresTwoFactorAuthError
	if errors.As(err, &err2fa) && slices.Contains(err2fa, "totp") {
		totpSecret := strings.ToUpper(strings.ReplaceAll(os.Getenv("VRC_TOTP_SECRET"), " ", ""))
		code, err := otp.TOTP(totpSecret)
		require.NoError(t, err)
		err = api.AuthAPI.Verify2FAWithTOTP(strconv.Itoa(int(code)))
		require.NoError(t, err)

		currentUser, err = api.AuthAPI.Login(username, password)
		require.NoError(t, err)
	} else {
		require.NoError(t, err)
	}
	authCookie, err := api.AuthCookie()
	require.NoError(t, err)
	require.NotEmpty(t, authCookie)
	t.Log(authCookie)

	t.Logf("Current user: %s", currentUser.DisplayName)
}

func TestAuthAPI_LoginWithAuthCookie(t *testing.T) {
	authCookie := os.Getenv("VRC_AUTH_COOKIE")
	if authCookie == "" {
		t.Skip("VRC_AUTH_COOKIE not set")
	}
	currentUser, err := api.AuthAPI.LoginWithAuthCookie(authCookie)
	require.NoError(t, err)

	t.Logf("Current user: %s", currentUser.DisplayName)
}

func TestAuthAPI_Verify2FAWithRecoveryCode(t *testing.T) {
	username, password := os.Getenv("VRC_USERNAME"), os.Getenv("VRC_PASSWORD")
	if username == "" || password == "" {
		t.Skip("VRC_USERNAME and VRC_PASSWORD not set")
	}

	recoveryCode := os.Getenv("VRC_RECOVERY_CODE")
	if recoveryCode == "" {
		t.Skip("VRC_RECOVERY_CODE not set")
	}

	currentUser, err := api.AuthAPI.Login(username, password)
	var err2fa RequiresTwoFactorAuthError
	require.ErrorAs(t, err, &err2fa)
	require.Contains(t, err2fa, "otp")
	err = api.AuthAPI.Verify2FAWithRecoveryCode(recoveryCode)
	require.NoError(t, err)

	currentUser, err = api.AuthAPI.Login(username, password)
	require.NoError(t, err)
	t.Logf("Current user: %s", currentUser.DisplayName)
}
