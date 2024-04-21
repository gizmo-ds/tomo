/*
 * Copyright (c) 2024 Gizmo.
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package vrcapi

import (
	"encoding/json"
	"errors"
	"fmt"

	"github.com/gizmo-ds/tomo/pkg/vrcapi/models"
	"github.com/go-resty/resty/v2"
)

type RequiresTwoFactorAuthError []string

func (e RequiresTwoFactorAuthError) Error() string {
	return fmt.Sprintf("requires two-factor authentication: %v", []string(e))
}

func handleAPIResponse(resp *resty.Response, expectedStatusCode int, err error, result *models.BaseResponse) error {
	if resp.StatusCode() == expectedStatusCode && err == nil {
		return nil
	}

	if result == nil {
		if err = json.Unmarshal(resp.Body(), &result); err != nil {
			return err
		}
	}
	var _errors []error
	if err != nil {
		_errors = append(_errors, err)
	}
	if result.Error != nil {
		_errors = append(_errors, errors.New(result.Error.Message))
	}
	if len(result.RequiresTwoFactorAuth) > 0 {
		_errors = append(_errors, RequiresTwoFactorAuthError(result.RequiresTwoFactorAuth))
	}
	return errors.Join(_errors...)
}
