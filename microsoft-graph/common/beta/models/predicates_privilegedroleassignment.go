package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedRoleAssignmentOperationPredicate struct {
	ExpirationDateTime *string
	Id                 *string
	IsElevated         *bool
	ODataType          *string
	ResultMessage      *string
	RoleId             *string
	UserId             *string
}

func (p PrivilegedRoleAssignmentOperationPredicate) Matches(input PrivilegedRoleAssignment) bool {

	if p.ExpirationDateTime != nil && (input.ExpirationDateTime == nil || *p.ExpirationDateTime != *input.ExpirationDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsElevated != nil && (input.IsElevated == nil || *p.IsElevated != *input.IsElevated) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ResultMessage != nil && (input.ResultMessage == nil || *p.ResultMessage != *input.ResultMessage) {
		return false
	}

	if p.RoleId != nil && (input.RoleId == nil || *p.RoleId != *input.RoleId) {
		return false
	}

	if p.UserId != nil && (input.UserId == nil || *p.UserId != *input.UserId) {
		return false
	}

	return true
}
