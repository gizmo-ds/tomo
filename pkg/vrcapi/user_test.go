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

	"github.com/stretchr/testify/require"
)

func TestUserAPI_Friends(t *testing.T) {
	authCookie := os.Getenv("VRC_AUTH_COOKIE")
	if authCookie == "" {
		t.Skip("VRC_AUTH_COOKIE not set")
	}
	currentUser, err := api.AuthAPI.LoginWithAuthCookie(authCookie)
	require.NoError(t, err)

	t.Logf("Current user: %s", currentUser.DisplayName)

	friends, err := api.UserAPI.Friends(-1, 0, false)
	require.NoError(t, err)
	require.NotEmpty(t, friends)
	t.Logf("Friends: %d", len(friends))
}

func TestUserAPI_Search(t *testing.T) {
	authCookie := os.Getenv("VRC_AUTH_COOKIE")
	if authCookie == "" {
		t.Skip("VRC_AUTH_COOKIE not set")
	}
	currentUser, err := api.AuthAPI.LoginWithAuthCookie(authCookie)
	require.NoError(t, err)

	t.Logf("Current user: %s", currentUser.DisplayName)

	users, err := api.UserAPI.Search("Gizmo", -1, 0)
	require.NoError(t, err)
	t.Log(users)
}
