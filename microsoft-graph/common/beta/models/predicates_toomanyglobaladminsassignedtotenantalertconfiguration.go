package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TooManyGlobalAdminsAssignedToTenantAlertConfigurationOperationPredicate struct {
	AlertDefinitionId                           *string
	GlobalAdminCountThreshold                   *int64
	Id                                          *string
	IsEnabled                                   *bool
	ODataType                                   *string
	PercentageOfGlobalAdminsOutOfRolesThreshold *int64
	ScopeId                                     *string
	ScopeType                                   *string
}

func (p TooManyGlobalAdminsAssignedToTenantAlertConfigurationOperationPredicate) Matches(input TooManyGlobalAdminsAssignedToTenantAlertConfiguration) bool {

	if p.AlertDefinitionId != nil && (input.AlertDefinitionId == nil || *p.AlertDefinitionId != *input.AlertDefinitionId) {
		return false
	}

	if p.GlobalAdminCountThreshold != nil && (input.GlobalAdminCountThreshold == nil || *p.GlobalAdminCountThreshold != *input.GlobalAdminCountThreshold) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsEnabled != nil && (input.IsEnabled == nil || *p.IsEnabled != *input.IsEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PercentageOfGlobalAdminsOutOfRolesThreshold != nil && (input.PercentageOfGlobalAdminsOutOfRolesThreshold == nil || *p.PercentageOfGlobalAdminsOutOfRolesThreshold != *input.PercentageOfGlobalAdminsOutOfRolesThreshold) {
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
