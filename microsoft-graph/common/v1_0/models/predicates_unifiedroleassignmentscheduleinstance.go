package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleAssignmentScheduleInstanceOperationPredicate struct {
	AppScopeId               *string
	AssignmentType           *string
	DirectoryScopeId         *string
	EndDateTime              *string
	Id                       *string
	MemberType               *string
	ODataType                *string
	PrincipalId              *string
	RoleAssignmentOriginId   *string
	RoleAssignmentScheduleId *string
	RoleDefinitionId         *string
	StartDateTime            *string
}

func (p UnifiedRoleAssignmentScheduleInstanceOperationPredicate) Matches(input UnifiedRoleAssignmentScheduleInstance) bool {

	if p.AppScopeId != nil && (input.AppScopeId == nil || *p.AppScopeId != *input.AppScopeId) {
		return false
	}

	if p.AssignmentType != nil && (input.AssignmentType == nil || *p.AssignmentType != *input.AssignmentType) {
		return false
	}

	if p.DirectoryScopeId != nil && (input.DirectoryScopeId == nil || *p.DirectoryScopeId != *input.DirectoryScopeId) {
		return false
	}

	if p.EndDateTime != nil && (input.EndDateTime == nil || *p.EndDateTime != *input.EndDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.MemberType != nil && (input.MemberType == nil || *p.MemberType != *input.MemberType) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PrincipalId != nil && (input.PrincipalId == nil || *p.PrincipalId != *input.PrincipalId) {
		return false
	}

	if p.RoleAssignmentOriginId != nil && (input.RoleAssignmentOriginId == nil || *p.RoleAssignmentOriginId != *input.RoleAssignmentOriginId) {
		return false
	}

	if p.RoleAssignmentScheduleId != nil && (input.RoleAssignmentScheduleId == nil || *p.RoleAssignmentScheduleId != *input.RoleAssignmentScheduleId) {
		return false
	}

	if p.RoleDefinitionId != nil && (input.RoleDefinitionId == nil || *p.RoleDefinitionId != *input.RoleDefinitionId) {
		return false
	}

	if p.StartDateTime != nil && (input.StartDateTime == nil || *p.StartDateTime != *input.StartDateTime) {
		return false
	}

	return true
}
