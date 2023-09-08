package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleManagementPolicyOperationPredicate struct {
	Description           *string
	DisplayName           *string
	Id                    *string
	IsOrganizationDefault *bool
	LastModifiedDateTime  *string
	ODataType             *string
	ScopeId               *string
	ScopeType             *string
}

func (p UnifiedRoleManagementPolicyOperationPredicate) Matches(input UnifiedRoleManagementPolicy) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsOrganizationDefault != nil && (input.IsOrganizationDefault == nil || *p.IsOrganizationDefault != *input.IsOrganizationDefault) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
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
