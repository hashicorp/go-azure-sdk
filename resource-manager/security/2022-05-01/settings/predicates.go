// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package settings

type SettingOperationPredicate struct {
}

func (p SettingOperationPredicate) Matches(input Setting) bool {

	return true
}
