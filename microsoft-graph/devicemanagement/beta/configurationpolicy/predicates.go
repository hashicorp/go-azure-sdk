package configurationpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

import "github.com/hashicorp/go-azure-sdk/microsoft-graph/common-types/beta"

type DeviceManagementConfigurationPolicyOperationPredicate struct {
}

func (p DeviceManagementConfigurationPolicyOperationPredicate) Matches(input beta.DeviceManagementConfigurationPolicy) bool {

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
