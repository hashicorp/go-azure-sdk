package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InboundSharedUserProfileOperationPredicate struct {
	DisplayName       *string
	HomeTenantId      *string
	ODataType         *string
	UserId            *string
	UserPrincipalName *string
}

func (p InboundSharedUserProfileOperationPredicate) Matches(input InboundSharedUserProfile) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.HomeTenantId != nil && (input.HomeTenantId == nil || *p.HomeTenantId != *input.HomeTenantId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UserId != nil && (input.UserId == nil || *p.UserId != *input.UserId) {
		return false
	}

	if p.UserPrincipalName != nil && (input.UserPrincipalName == nil || *p.UserPrincipalName != *input.UserPrincipalName) {
		return false
	}

	return true
}
