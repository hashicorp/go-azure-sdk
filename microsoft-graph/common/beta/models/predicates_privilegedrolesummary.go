package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedRoleSummaryOperationPredicate struct {
	ElevatedCount *int64
	Id            *string
	ManagedCount  *int64
	MfaEnabled    *bool
	ODataType     *string
	UsersCount    *int64
}

func (p PrivilegedRoleSummaryOperationPredicate) Matches(input PrivilegedRoleSummary) bool {

	if p.ElevatedCount != nil && (input.ElevatedCount == nil || *p.ElevatedCount != *input.ElevatedCount) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ManagedCount != nil && (input.ManagedCount == nil || *p.ManagedCount != *input.ManagedCount) {
		return false
	}

	if p.MfaEnabled != nil && (input.MfaEnabled == nil || *p.MfaEnabled != *input.MfaEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UsersCount != nil && (input.UsersCount == nil || *p.UsersCount != *input.UsersCount) {
		return false
	}

	return true
}
