package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AttackSimulationOperationOperationPredicate struct {
	CreatedDateTime     *string
	Id                  *string
	LastActionDateTime  *string
	ODataType           *string
	PercentageCompleted *int64
	ResourceLocation    *string
	StatusDetail        *string
	TenantId            *string
}

func (p AttackSimulationOperationOperationPredicate) Matches(input AttackSimulationOperation) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastActionDateTime != nil && (input.LastActionDateTime == nil || *p.LastActionDateTime != *input.LastActionDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PercentageCompleted != nil && (input.PercentageCompleted == nil || *p.PercentageCompleted != *input.PercentageCompleted) {
		return false
	}

	if p.ResourceLocation != nil && (input.ResourceLocation == nil || *p.ResourceLocation != *input.ResourceLocation) {
		return false
	}

	if p.StatusDetail != nil && (input.StatusDetail == nil || *p.StatusDetail != *input.StatusDetail) {
		return false
	}

	if p.TenantId != nil && (input.TenantId == nil || *p.TenantId != *input.TenantId) {
		return false
	}

	return true
}
