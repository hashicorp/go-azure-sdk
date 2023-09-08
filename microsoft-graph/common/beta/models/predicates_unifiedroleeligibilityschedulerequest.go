package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleEligibilityScheduleRequestOperationPredicate struct {
	Action            *string
	AppScopeId        *string
	ApprovalId        *string
	CompletedDateTime *string
	CreatedDateTime   *string
	CustomData        *string
	DirectoryScopeId  *string
	Id                *string
	IsValidationOnly  *bool
	Justification     *string
	ODataType         *string
	PrincipalId       *string
	RoleDefinitionId  *string
	Status            *string
	TargetScheduleId  *string
}

func (p UnifiedRoleEligibilityScheduleRequestOperationPredicate) Matches(input UnifiedRoleEligibilityScheduleRequest) bool {

	if p.Action != nil && (input.Action == nil || *p.Action != *input.Action) {
		return false
	}

	if p.AppScopeId != nil && (input.AppScopeId == nil || *p.AppScopeId != *input.AppScopeId) {
		return false
	}

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

	if p.DirectoryScopeId != nil && (input.DirectoryScopeId == nil || *p.DirectoryScopeId != *input.DirectoryScopeId) {
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

	if p.RoleDefinitionId != nil && (input.RoleDefinitionId == nil || *p.RoleDefinitionId != *input.RoleDefinitionId) {
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
