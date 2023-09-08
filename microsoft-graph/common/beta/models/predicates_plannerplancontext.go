package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanContextOperationPredicate struct {
	AssociationType   *string
	CreatedDateTime   *string
	IsCreationContext *bool
	ODataType         *string
	OwnerAppId        *string
}

func (p PlannerPlanContextOperationPredicate) Matches(input PlannerPlanContext) bool {

	if p.AssociationType != nil && (input.AssociationType == nil || *p.AssociationType != *input.AssociationType) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.IsCreationContext != nil && (input.IsCreationContext == nil || *p.IsCreationContext != *input.IsCreationContext) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OwnerAppId != nil && (input.OwnerAppId == nil || *p.OwnerAppId != *input.OwnerAppId) {
		return false
	}

	return true
}
