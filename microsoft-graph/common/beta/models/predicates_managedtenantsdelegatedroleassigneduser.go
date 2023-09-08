package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsDelegatedRoleAssignedUserOperationPredicate struct {
	DisplayName       *string
	ODataType         *string
	UserEntityId      *string
	UserPrincipalName *string
}

func (p ManagedTenantsDelegatedRoleAssignedUserOperationPredicate) Matches(input ManagedTenantsDelegatedRoleAssignedUser) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UserEntityId != nil && (input.UserEntityId == nil || *p.UserEntityId != *input.UserEntityId) {
		return false
	}

	if p.UserPrincipalName != nil && (input.UserPrincipalName == nil || *p.UserPrincipalName != *input.UserPrincipalName) {
		return false
	}

	return true
}
