package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupEligibilityScheduleOperationPredicate struct {
	CreatedDateTime  *string
	CreatedUsing     *string
	GroupId          *string
	Id               *string
	ModifiedDateTime *string
	ODataType        *string
	PrincipalId      *string
	Status           *string
}

func (p PrivilegedAccessGroupEligibilityScheduleOperationPredicate) Matches(input PrivilegedAccessGroupEligibilitySchedule) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.CreatedUsing != nil && (input.CreatedUsing == nil || *p.CreatedUsing != *input.CreatedUsing) {
		return false
	}

	if p.GroupId != nil && (input.GroupId == nil || *p.GroupId != *input.GroupId) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ModifiedDateTime != nil && (input.ModifiedDateTime == nil || *p.ModifiedDateTime != *input.ModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PrincipalId != nil && (input.PrincipalId == nil || *p.PrincipalId != *input.PrincipalId) {
		return false
	}

	if p.Status != nil && (input.Status == nil || *p.Status != *input.Status) {
		return false
	}

	return true
}
