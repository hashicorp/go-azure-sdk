package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ShiftOperationPredicate struct {
	CreatedDateTime      *string
	Id                   *string
	LastModifiedDateTime *string
	ODataType            *string
	SchedulingGroupId    *string
	UserId               *string
}

func (p ShiftOperationPredicate) Matches(input Shift) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SchedulingGroupId != nil && (input.SchedulingGroupId == nil || *p.SchedulingGroupId != *input.SchedulingGroupId) {
		return false
	}

	if p.UserId != nil && (input.UserId == nil || *p.UserId != *input.UserId) {
		return false
	}

	return true
}
