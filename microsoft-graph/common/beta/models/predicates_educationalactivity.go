package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EducationalActivityOperationPredicate struct {
	CompletionMonthYear  *string
	CreatedDateTime      *string
	EndMonthYear         *string
	Id                   *string
	IsSearchable         *bool
	LastModifiedDateTime *string
	ODataType            *string
	StartMonthYear       *string
}

func (p EducationalActivityOperationPredicate) Matches(input EducationalActivity) bool {

	if p.CompletionMonthYear != nil && (input.CompletionMonthYear == nil || *p.CompletionMonthYear != *input.CompletionMonthYear) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.EndMonthYear != nil && (input.EndMonthYear == nil || *p.EndMonthYear != *input.EndMonthYear) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsSearchable != nil && (input.IsSearchable == nil || *p.IsSearchable != *input.IsSearchable) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.StartMonthYear != nil && (input.StartMonthYear == nil || *p.StartMonthYear != *input.StartMonthYear) {
		return false
	}

	return true
}
