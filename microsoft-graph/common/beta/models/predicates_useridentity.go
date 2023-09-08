package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserIdentityOperationPredicate struct {
	DisplayName       *string
	Id                *string
	IpAddress         *string
	ODataType         *string
	UserPrincipalName *string
}

func (p UserIdentityOperationPredicate) Matches(input UserIdentity) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IpAddress != nil && (input.IpAddress == nil || *p.IpAddress != *input.IpAddress) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UserPrincipalName != nil && (input.UserPrincipalName == nil || *p.UserPrincipalName != *input.UserPrincipalName) {
		return false
	}

	return true
}
