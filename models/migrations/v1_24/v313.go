// Copyright 2024 The Gitea Authors. All rights reserved.
// SPDX-License-Identifier: MIT

package v1_24 //nolint

import (
	"xorm.io/xorm"
)

func AddProtectedToSecrets(x *xorm.Engine) error {
	type secrets struct {
		Protected bool
	}

	return x.Sync(new(secrets))
}
