package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OpenShiftOperationPredicate struct {
	CreatedDateTime      *string
	Id                   *string
	LastModifiedDateTime *string
	ODataType            *string
	SchedulingGroupId    *string
}

func (p OpenShiftOperationPredicate) Matches(input OpenShift) bool {

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

	return true
}
