package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteOperationOperationPredicate struct {
	CreatedDateTime    *string
	Id                 *string
	LastActionDateTime *string
	ODataType          *string
	PercentComplete    *string
	ResourceId         *string
	ResourceLocation   *string
}

func (p OnenoteOperationOperationPredicate) Matches(input OnenoteOperation) bool {

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

	if p.PercentComplete != nil && (input.PercentComplete == nil || *p.PercentComplete != *input.PercentComplete) {
		return false
	}

	if p.ResourceId != nil && (input.ResourceId == nil || *p.ResourceId != *input.ResourceId) {
		return false
	}

	if p.ResourceLocation != nil && (input.ResourceLocation == nil || *p.ResourceLocation != *input.ResourceLocation) {
		return false
	}

	return true
}
