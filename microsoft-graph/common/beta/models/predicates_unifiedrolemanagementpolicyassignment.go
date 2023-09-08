package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleManagementPolicyAssignmentOperationPredicate struct {
	Id               *string
	ODataType        *string
	PolicyId         *string
	RoleDefinitionId *string
	ScopeId          *string
	ScopeType        *string
}

func (p UnifiedRoleManagementPolicyAssignmentOperationPredicate) Matches(input UnifiedRoleManagementPolicyAssignment) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PolicyId != nil && (input.PolicyId == nil || *p.PolicyId != *input.PolicyId) {
		return false
	}

	if p.RoleDefinitionId != nil && (input.RoleDefinitionId == nil || *p.RoleDefinitionId != *input.RoleDefinitionId) {
		return false
	}

	if p.ScopeId != nil && (input.ScopeId == nil || *p.ScopeId != *input.ScopeId) {
		return false
	}

	if p.ScopeType != nil && (input.ScopeType == nil || *p.ScopeType != *input.ScopeType) {
		return false
	}

	return true
}
