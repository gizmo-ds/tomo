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
	"net/http"
	"time"

	"github.com/gizmo-ds/tomo/pkg/vrcapi/models"
)

var (
	ErrLoginFail               = errors.New("login fail")
	ErrVerifyTwoFactorAuthFail = errors.New("verify two-factor authentication fail")
)

type AuthAPI interface {
	// Login authenticates a user using username and password and sets an 'auth' cookie to maintain the session.
	//
	//	Note: Be mindful of the API's session limit, as excessive logins may lead to denial of access.
	Login(username, password string) (models.CurrentUser, error)

	// LoginWithAuthCookie authenticates a user using a previously set 'auth' cookie.
	// It is useful for systems that maintain long-lived sessions or where minimizing re-authentication is desired.
	LoginWithAuthCookie(cookie string) (models.CurrentUser, error)

	// VerifyTwoFactorAuthTOTP verifies the TOTP code provided for two-factor authentication.
	VerifyTwoFactorAuthTOTP(code string) error

	// VerifyTwoFactorAuthEmail verifies the email verification code provided for two-factor authentication.
	VerifyTwoFactorAuthEmail(code string) error

	// Logout terminates the user's session and clears any session-related data.
	Logout() error
}

type authAPI struct {
	*common
}

func (a *authAPI) Login(username, password string) (models.CurrentUser, error) {
	var result models.AuthResponse
	resp, err := a.httpClient.R().
		SetBasicAuth(username, password).
		SetResult(&result).
		Get("/auth/user")
	if err != nil {
		return result.CurrentUser, err
	}
	return result.CurrentUser, handleAPIResponse(resp, http.StatusOK, ErrLoginFail, result.BaseResponse)
}

func (a *authAPI) LoginWithAuthCookie(authCookie string) (models.CurrentUser, error) {
	var result models.AuthResponse

	a.httpClient.SetCookie(&http.Cookie{Name: "auth", Value: authCookie})
	resp, err := a.httpClient.R().
		SetResult(&result).
		Get("/auth/user")
	if err != nil {
		return result.CurrentUser, err
	}
	return result.CurrentUser, handleAPIResponse(resp, http.StatusOK, ErrLoginFail, result.BaseResponse)
}

func (a *authAPI) VerifyTwoFactorAuthTOTP(code string) error {
	var result models.VerifyTwoFactorAuthResponse
	resp, err := a.httpClient.R().
		SetBody(map[string]any{"code": code}).
		SetResult(&result).
		Post("/auth/twofactorauth/totp/verify")
	if err != nil {
		return err
	}
	if err = handleAPIResponse(resp, http.StatusOK, ErrVerifyTwoFactorAuthFail, result.BaseResponse); err != nil {
		return err
	}
	if !result.Verified {
		return ErrVerifyTwoFactorAuthFail
	}
	return nil
}

func (a *authAPI) VerifyTwoFactorAuthEmail(code string) error {
	var result models.VerifyTwoFactorAuthResponse
	resp, err := a.httpClient.R().
		SetBody(map[string]any{"code": code}).
		SetResult(&result).
		Post("/auth/twofactorauth/emailotp/verify")
	if err != nil {
		return err
	}
	if err = handleAPIResponse(resp, http.StatusOK, ErrVerifyTwoFactorAuthFail, result.BaseResponse); err != nil {
		return err
	}
	if !result.Verified {
		return ErrVerifyTwoFactorAuthFail
	}
	return nil
}

func (a *authAPI) Logout() error {
	resp, err := a.httpClient.R().Put("/logout")
	if err != nil {
		return err
	}
	err = handleAPIResponse(resp, http.StatusOK, nil, nil)
	if err != nil {
		return err
	}
	a.httpClient.SetCookie(
		&http.Cookie{Name: "auth", Value: "", MaxAge: -1, Expires: time.Unix(1, 0)})
	return nil
}
