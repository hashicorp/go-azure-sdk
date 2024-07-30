package compliancepolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type DeviceManagementCompliancePolicyOperationPredicate struct {
}

func (p DeviceManagementCompliancePolicyOperationPredicate) Matches(input beta.DeviceManagementCompliancePolicy) bool {

	return true
}

type DeviceManagementComplianceScheduledActionForRuleOperationPredicate struct {
	Id        *string
	ODataType *string
	RuleName  *string
}

func (p DeviceManagementComplianceScheduledActionForRuleOperationPredicate) Matches(input beta.DeviceManagementComplianceScheduledActionForRule) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RuleName != nil && (input.RuleName == nil || *p.RuleName != *input.RuleName) {
		return false
	}

	return true
}

type DeviceManagementConfigurationPolicyAssignmentOperationPredicate struct {
	Id        *string
	ODataType *string
	SourceId  *string
}

func (p DeviceManagementConfigurationPolicyAssignmentOperationPredicate) Matches(input beta.DeviceManagementConfigurationPolicyAssignment) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SourceId != nil && (input.SourceId == nil || *p.SourceId != *input.SourceId) {
		return false
	}

	return true
}
