package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProvisionedPlanOperationPredicate struct {
	CapabilityStatus   *string
	ODataType          *string
	ProvisioningStatus *string
	Service            *string
}

func (p ProvisionedPlanOperationPredicate) Matches(input ProvisionedPlan) bool {

	if p.CapabilityStatus != nil && (input.CapabilityStatus == nil || *p.CapabilityStatus != *input.CapabilityStatus) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ProvisioningStatus != nil && (input.ProvisioningStatus == nil || *p.ProvisioningStatus != *input.ProvisioningStatus) {
		return false
	}

	if p.Service != nil && (input.Service == nil || *p.Service != *input.Service) {
		return false
	}

	return true
}
