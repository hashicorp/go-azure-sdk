package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlannerPlanConfigurationOperationPredicate struct {
	CreatedDateTime      *string
	DefaultLanguage      *string
	Id                   *string
	LastModifiedDateTime *string
	ODataType            *string
}

func (p PlannerPlanConfigurationOperationPredicate) Matches(input PlannerPlanConfiguration) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DefaultLanguage != nil && (input.DefaultLanguage == nil || *p.DefaultLanguage != *input.DefaultLanguage) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
