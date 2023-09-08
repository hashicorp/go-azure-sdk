package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivilegedAccessGroupAssignmentScheduleRequestOperationPredicate struct {
	ApprovalId        *string
	CompletedDateTime *string
	CreatedDateTime   *string
	CustomData        *string
	GroupId           *string
	Id                *string
	IsValidationOnly  *bool
	Justification     *string
	ODataType         *string
	PrincipalId       *string
	Status            *string
	TargetScheduleId  *string
}

func (p PrivilegedAccessGroupAssignmentScheduleRequestOperationPredicate) Matches(input PrivilegedAccessGroupAssignmentScheduleRequest) bool {

	if p.ApprovalId != nil && (input.ApprovalId == nil || *p.ApprovalId != *input.ApprovalId) {
		return false
	}

	if p.CompletedDateTime != nil && (input.CompletedDateTime == nil || *p.CompletedDateTime != *input.CompletedDateTime) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.CustomData != nil && (input.CustomData == nil || *p.CustomData != *input.CustomData) {
		return false
	}

	if p.GroupId != nil && (input.GroupId == nil || *p.GroupId != *input.GroupId) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsValidationOnly != nil && (input.IsValidationOnly == nil || *p.IsValidationOnly != *input.IsValidationOnly) {
		return false
	}

	if p.Justification != nil && (input.Justification == nil || *p.Justification != *input.Justification) {
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

	if p.TargetScheduleId != nil && (input.TargetScheduleId == nil || *p.TargetScheduleId != *input.TargetScheduleId) {
		return false
	}

	return true
}
