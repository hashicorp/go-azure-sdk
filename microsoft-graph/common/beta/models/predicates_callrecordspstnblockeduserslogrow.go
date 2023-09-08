package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CallRecordsPstnBlockedUsersLogRowOperationPredicate struct {
	BlockDateTime       *string
	BlockReason         *string
	ODataType           *string
	RemediationId       *string
	UserDisplayName     *string
	UserId              *string
	UserPrincipalName   *string
	UserTelephoneNumber *string
}

func (p CallRecordsPstnBlockedUsersLogRowOperationPredicate) Matches(input CallRecordsPstnBlockedUsersLogRow) bool {

	if p.BlockDateTime != nil && (input.BlockDateTime == nil || *p.BlockDateTime != *input.BlockDateTime) {
		return false
	}

	if p.BlockReason != nil && (input.BlockReason == nil || *p.BlockReason != *input.BlockReason) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RemediationId != nil && (input.RemediationId == nil || *p.RemediationId != *input.RemediationId) {
		return false
	}

	if p.UserDisplayName != nil && (input.UserDisplayName == nil || *p.UserDisplayName != *input.UserDisplayName) {
		return false
	}

	if p.UserId != nil && (input.UserId == nil || *p.UserId != *input.UserId) {
		return false
	}

	if p.UserPrincipalName != nil && (input.UserPrincipalName == nil || *p.UserPrincipalName != *input.UserPrincipalName) {
		return false
	}

	if p.UserTelephoneNumber != nil && (input.UserTelephoneNumber == nil || *p.UserTelephoneNumber != *input.UserTelephoneNumber) {
		return false
	}

	return true
}
