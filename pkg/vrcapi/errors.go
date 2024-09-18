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
	"go.uber.org/multierr"
)

type RequiresTwoFactorAuthError []string

func (e RequiresTwoFactorAuthError) Error() string {
	return fmt.Sprintf("requires two-factor authentication: %v", []string(e))
}

func handleAPIResponse(resp *resty.Response, expectedStatusCode int, failedError error, result *models.BaseResponse) error {
	if result == nil {
		_ = json.Unmarshal(resp.Body(), &result)
	}
	var err error
	if result != nil && result.Error != nil {
		err = multierr.Append(err, errors.New(result.Error.Message))
	}
	if result != nil && len(result.RequiresTwoFactorAuth) > 0 {
		err = multierr.Append(err, RequiresTwoFactorAuthError(result.RequiresTwoFactorAuth))
	}
	if resp.StatusCode() != expectedStatusCode || failedError == nil {
		err = multierr.Append(err, failedError)
	}
	return err
}
