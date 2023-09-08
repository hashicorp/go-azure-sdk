package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignedPlanOperationPredicate struct {
	AssignedDateTime *string
	CapabilityStatus *string
	ODataType        *string
	Service          *string
	ServicePlanId    *string
}

func (p AssignedPlanOperationPredicate) Matches(input AssignedPlan) bool {

	if p.AssignedDateTime != nil && (input.AssignedDateTime == nil || *p.AssignedDateTime != *input.AssignedDateTime) {
		return false
	}

	if p.CapabilityStatus != nil && (input.CapabilityStatus == nil || *p.CapabilityStatus != *input.CapabilityStatus) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Service != nil && (input.Service == nil || *p.Service != *input.Service) {
		return false
	}

	if p.ServicePlanId != nil && (input.ServicePlanId == nil || *p.ServicePlanId != *input.ServicePlanId) {
		return false
	}

	return true
}
