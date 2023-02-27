// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package roles

type RoleOperationPredicate struct {
}

func (p RoleOperationPredicate) Matches(input Role) bool {

	return true
}
