package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsWorkloadActionDeploymentStatusOperationPredicate struct {
	ActionId               *string
	DeployedPolicyId       *string
	IncludeAllUsers        *bool
	LastDeploymentDateTime *string
	ODataType              *string
}

func (p ManagedTenantsWorkloadActionDeploymentStatusOperationPredicate) Matches(input ManagedTenantsWorkloadActionDeploymentStatus) bool {

	if p.ActionId != nil && (input.ActionId == nil || *p.ActionId != *input.ActionId) {
		return false
	}

	if p.DeployedPolicyId != nil && (input.DeployedPolicyId == nil || *p.DeployedPolicyId != *input.DeployedPolicyId) {
		return false
	}

	if p.IncludeAllUsers != nil && (input.IncludeAllUsers == nil || *p.IncludeAllUsers != *input.IncludeAllUsers) {
		return false
	}

	if p.LastDeploymentDateTime != nil && (input.LastDeploymentDateTime == nil || *p.LastDeploymentDateTime != *input.LastDeploymentDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
