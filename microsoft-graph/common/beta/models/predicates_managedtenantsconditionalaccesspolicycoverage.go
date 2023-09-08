package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsConditionalAccessPolicyCoverageOperationPredicate struct {
	ConditionalAccessPolicyState *string
	Id                           *string
	LatestPolicyModifiedDateTime *string
	ODataType                    *string
	RequiresDeviceCompliance     *bool
	TenantDisplayName            *string
}

func (p ManagedTenantsConditionalAccessPolicyCoverageOperationPredicate) Matches(input ManagedTenantsConditionalAccessPolicyCoverage) bool {

	if p.ConditionalAccessPolicyState != nil && (input.ConditionalAccessPolicyState == nil || *p.ConditionalAccessPolicyState != *input.ConditionalAccessPolicyState) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LatestPolicyModifiedDateTime != nil && (input.LatestPolicyModifiedDateTime == nil || *p.LatestPolicyModifiedDateTime != *input.LatestPolicyModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RequiresDeviceCompliance != nil && (input.RequiresDeviceCompliance == nil || *p.RequiresDeviceCompliance != *input.RequiresDeviceCompliance) {
		return false
	}

	if p.TenantDisplayName != nil && (input.TenantDisplayName == nil || *p.TenantDisplayName != *input.TenantDisplayName) {
		return false
	}

	return true
}
