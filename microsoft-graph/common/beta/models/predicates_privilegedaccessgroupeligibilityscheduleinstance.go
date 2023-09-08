package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupEligibilityScheduleInstanceOperationPredicate struct {
	EligibilityScheduleId *string
	EndDateTime           *string
	GroupId               *string
	Id                    *string
	ODataType             *string
	PrincipalId           *string
	StartDateTime         *string
}

func (p PrivilegedAccessGroupEligibilityScheduleInstanceOperationPredicate) Matches(input PrivilegedAccessGroupEligibilityScheduleInstance) bool {

	if p.EligibilityScheduleId != nil && (input.EligibilityScheduleId == nil || *p.EligibilityScheduleId != *input.EligibilityScheduleId) {
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
