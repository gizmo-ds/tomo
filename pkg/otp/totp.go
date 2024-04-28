/*
 * Copyright (c) 2024 Gizmo.
 *
 * This Source Code Form is subject to the terms of the Mozilla Public
 * License, v. 2.0. If a copy of the MPL was not distributed with this
 * file, You can obtain one at https://mozilla.org/MPL/2.0/.
 */

package otp

import (
	"crypto/hmac"
	"crypto/sha1"
	"encoding/base32"
	"encoding/binary"
	"time"
)

func TOTP(secret string, t ...int64) (uint32, error) {
	_t := time.Now().Unix() / 30
	if len(t) > 0 {
		_t = t[0] / 30
	}
	k, err := base32.StdEncoding.DecodeString(secret)
	if err != nil {
		return 0, err
	}

	hash := hmac.New(sha1.New, k)
	err = binary.Write(hash, binary.BigEndian, _t)
	if err != nil {
		return 0, err
	}
	h := hash.Sum(nil)

	offset := h[19] & 0x0f

	result := binary.BigEndian.Uint32(h[offset : offset+4])

	result &= 0x7fffffff
	code := result % 1000000

	return code, nil
}
