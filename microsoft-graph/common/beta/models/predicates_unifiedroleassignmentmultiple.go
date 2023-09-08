package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleAssignmentMultipleOperationPredicate struct {
	Condition        *string
	Description      *string
	DisplayName      *string
	Id               *string
	ODataType        *string
	RoleDefinitionId *string
}

func (p UnifiedRoleAssignmentMultipleOperationPredicate) Matches(input UnifiedRoleAssignmentMultiple) bool {

	if p.Condition != nil && (input.Condition == nil || *p.Condition != *input.Condition) {
		return false
	}

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RoleDefinitionId != nil && (input.RoleDefinitionId == nil || *p.RoleDefinitionId != *input.RoleDefinitionId) {
		return false
	}

	return true
}
