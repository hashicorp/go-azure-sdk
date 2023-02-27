// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package useridentity

type UserIdentityContractOperationPredicate struct {
	Id       *string
	Provider *string
}

func (p UserIdentityContractOperationPredicate) Matches(input UserIdentityContract) bool {

	if p.Id != nil && (input.Id == nil && *p.Id != *input.Id) {
		return false
	}

	if p.Provider != nil && (input.Provider == nil && *p.Provider != *input.Provider) {
		return false
	}

	return true
}
