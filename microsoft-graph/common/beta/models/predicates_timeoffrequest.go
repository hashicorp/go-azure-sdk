package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TimeOffRequestOperationPredicate struct {
	CreatedDateTime       *string
	EndDateTime           *string
	Id                    *string
	LastModifiedDateTime  *string
	ManagerActionDateTime *string
	ManagerActionMessage  *string
	ManagerUserId         *string
	ODataType             *string
	SenderDateTime        *string
	SenderMessage         *string
	SenderUserId          *string
	StartDateTime         *string
	TimeOffReasonId       *string
}

func (p TimeOffRequestOperationPredicate) Matches(input TimeOffRequest) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.EndDateTime != nil && (input.EndDateTime == nil || *p.EndDateTime != *input.EndDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ManagerActionDateTime != nil && (input.ManagerActionDateTime == nil || *p.ManagerActionDateTime != *input.ManagerActionDateTime) {
		return false
	}

	if p.ManagerActionMessage != nil && (input.ManagerActionMessage == nil || *p.ManagerActionMessage != *input.ManagerActionMessage) {
		return false
	}

	if p.ManagerUserId != nil && (input.ManagerUserId == nil || *p.ManagerUserId != *input.ManagerUserId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SenderDateTime != nil && (input.SenderDateTime == nil || *p.SenderDateTime != *input.SenderDateTime) {
		return false
	}

	if p.SenderMessage != nil && (input.SenderMessage == nil || *p.SenderMessage != *input.SenderMessage) {
		return false
	}

	if p.SenderUserId != nil && (input.SenderUserId == nil || *p.SenderUserId != *input.SenderUserId) {
		return false
	}

	if p.StartDateTime != nil && (input.StartDateTime == nil || *p.StartDateTime != *input.StartDateTime) {
		return false
	}

	if p.TimeOffReasonId != nil && (input.TimeOffReasonId == nil || *p.TimeOffReasonId != *input.TimeOffReasonId) {
		return false
	}

	return true
}
