package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UnifiedRoleManagementAlertDefinitionOperationPredicate struct {
	Description     *string
	DisplayName     *string
	HowToPrevent    *string
	Id              *string
	IsConfigurable  *bool
	IsRemediatable  *bool
	MitigationSteps *string
	ODataType       *string
	ScopeId         *string
	ScopeType       *string
	SecurityImpact  *string
}

func (p UnifiedRoleManagementAlertDefinitionOperationPredicate) Matches(input UnifiedRoleManagementAlertDefinition) bool {

	if p.Description != nil && (input.Description == nil || *p.Description != *input.Description) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.HowToPrevent != nil && (input.HowToPrevent == nil || *p.HowToPrevent != *input.HowToPrevent) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsConfigurable != nil && (input.IsConfigurable == nil || *p.IsConfigurable != *input.IsConfigurable) {
		return false
	}

	if p.IsRemediatable != nil && (input.IsRemediatable == nil || *p.IsRemediatable != *input.IsRemediatable) {
		return false
	}

	if p.MitigationSteps != nil && (input.MitigationSteps == nil || *p.MitigationSteps != *input.MitigationSteps) {
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

	if p.SecurityImpact != nil && (input.SecurityImpact == nil || *p.SecurityImpact != *input.SecurityImpact) {
		return false
	}

	return true
}
