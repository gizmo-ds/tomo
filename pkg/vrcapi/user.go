/*
 * Copyright (c) 2024 Gizmo.
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package vrcapi

import (
	"math"
	"net/http"
	"strconv"

	"github.com/duke-git/lancet/v2/condition"
	"github.com/gizmo-ds/tomo/pkg/utils"
	"github.com/gizmo-ds/tomo/pkg/vrcapi/models"
)

type UserAPI interface {
	Friends(n, offset int, offline bool) ([]models.Friend, error)
	Search(keyword string, n, offset int) ([]models.SearchUser, error)
	CurrentUser() (models.CurrentUser, error)
}

type userAPI struct {
	*common
}

func (u *userAPI) CurrentUser() (models.CurrentUser, error) {
	var result models.CurrentUserResponse
	resp, err := u.httpClient.R().
		SetResult(&result).
		Get("/auth/user")
	if err != nil {
		return result.CurrentUser, err
	}
	return result.CurrentUser, handleAPIResponse(resp, http.StatusOK, ErrRequestFailed, result.BaseResponse)
}

func (u *userAPI) Friends(n, offset int, offline bool) ([]models.Friend, error) {
	var result []models.Friend
	req := u.httpClient.R().
		SetResult(&result).
		SetQueryParam("n", strconv.Itoa(utils.Clamp(n, 1, 100, 60))).
		SetQueryParam("offset", strconv.Itoa(utils.Clamp(offset, 0, math.MaxInt, 0))).
		SetQueryParam("offline", condition.TernaryOperator(offline, "true", "false"))
	if offline {
		req.SetQueryParam("offline", "true")
	}
	resp, err := req.Get("/auth/user/friends")
	if err != nil {
		return nil, err
	}
	if err = handleAPIResponse(resp, http.StatusOK, ErrRequestFailed, nil); err != nil {
		return nil, err
	}
	return result, nil
}

func (u *userAPI) Search(keyword string, n, offset int) ([]models.SearchUser, error) {
	var result []models.SearchUser
	resp, err := u.httpClient.R().
		SetResult(&result).
		SetQueryParam("search", keyword).
		SetQueryParam("n", strconv.Itoa(utils.Clamp(n, 1, 100, 60))).
		SetQueryParam("offset", strconv.Itoa(utils.Clamp(offset, 0, math.MaxInt, 0))).
		Get("/users")
	if err != nil {
		return nil, err
	}
	if err = handleAPIResponse(resp, http.StatusOK, ErrRequestFailed, nil); err != nil {
		return nil, err
	}
	return result, err
}
