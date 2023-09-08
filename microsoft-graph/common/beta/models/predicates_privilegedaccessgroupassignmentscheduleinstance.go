package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupAssignmentScheduleInstanceOperationPredicate struct {
	AssignmentScheduleId *string
	EndDateTime          *string
	GroupId              *string
	Id                   *string
	ODataType            *string
	PrincipalId          *string
	StartDateTime        *string
}

func (p PrivilegedAccessGroupAssignmentScheduleInstanceOperationPredicate) Matches(input PrivilegedAccessGroupAssignmentScheduleInstance) bool {

	if p.AssignmentScheduleId != nil && (input.AssignmentScheduleId == nil || *p.AssignmentScheduleId != *input.AssignmentScheduleId) {
		return false
	}

	if p.EndDateTime != nil && (input.EndDateTime == nil || *p.EndDateTime != *input.EndDateTime) {
		return false
	}

	if p.GroupId != nil && (input.GroupId == nil || *p.GroupId != *input.GroupId) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PrincipalId != nil && (input.PrincipalId == nil || *p.PrincipalId != *input.PrincipalId) {
		return false
	}

	if p.StartDateTime != nil && (input.StartDateTime == nil || *p.StartDateTime != *input.StartDateTime) {
		return false
	}

	return true
}
