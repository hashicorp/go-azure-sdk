package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AadUserConversationMemberOperationPredicate struct {
	DisplayName                 *string
	Email                       *string
	Id                          *string
	ODataType                   *string
	TenantId                    *string
	UserId                      *string
	VisibleHistoryStartDateTime *string
}

func (p AadUserConversationMemberOperationPredicate) Matches(input AadUserConversationMember) bool {

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Email != nil && (input.Email == nil || *p.Email != *input.Email) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TenantId != nil && (input.TenantId == nil || *p.TenantId != *input.TenantId) {
		return false
	}

	if p.UserId != nil && (input.UserId == nil || *p.UserId != *input.UserId) {
		return false
	}

	if p.VisibleHistoryStartDateTime != nil && (input.VisibleHistoryStartDateTime == nil || *p.VisibleHistoryStartDateTime != *input.VisibleHistoryStartDateTime) {
		return false
	}

	return true
}
