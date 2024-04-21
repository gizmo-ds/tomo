/*
 * Copyright (c) 2024 Gizmo.
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package models

type (
	BaseResponse struct {
		Success               *SuccessResponse `json:"success,omitempty"`
		Error                 *ErrorResponse   `json:"error,omitempty"`
		RequiresTwoFactorAuth []string         `json:"requiresTwoFactorAuth,omitempty"`
	}
	SuccessResponse struct {
		Message    string `json:"message"`
		StatusCode int    `json:"status_code"`
	}
	ErrorResponse = SuccessResponse
)
