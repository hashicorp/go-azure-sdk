package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleAssignmentOperationPredicate struct {
	AppScopeId       *string
	Condition        *string
	DirectoryScopeId *string
	Id               *string
	ODataType        *string
	PrincipalId      *string
	RoleDefinitionId *string
}

func (p UnifiedRoleAssignmentOperationPredicate) Matches(input UnifiedRoleAssignment) bool {

	if p.AppScopeId != nil && (input.AppScopeId == nil || *p.AppScopeId != *input.AppScopeId) {
		return false
	}

	if p.Condition != nil && (input.Condition == nil || *p.Condition != *input.Condition) {
		return false
	}

	if p.DirectoryScopeId != nil && (input.DirectoryScopeId == nil || *p.DirectoryScopeId != *input.DirectoryScopeId) {
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

	if p.RoleDefinitionId != nil && (input.RoleDefinitionId == nil || *p.RoleDefinitionId != *input.RoleDefinitionId) {
		return false
	}

	return true
}
