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
}

func (p DeviceManagementComplianceScheduledActionForRuleOperationPredicate) Matches(input beta.DeviceManagementComplianceScheduledActionForRule) bool {

	return true
}

type DeviceManagementConfigurationPolicyAssignmentOperationPredicate struct {
}

func (p DeviceManagementConfigurationPolicyAssignmentOperationPredicate) Matches(input beta.DeviceManagementConfigurationPolicyAssignment) bool {

	return true
}
