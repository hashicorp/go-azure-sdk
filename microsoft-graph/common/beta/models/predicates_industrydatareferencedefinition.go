package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IndustryDataReferenceDefinitionOperationPredicate struct {
	Code                 *string
	CreatedDateTime      *string
	Id                   *string
	IsDisabled           *bool
	LastModifiedDateTime *string
	ODataType            *string
	ReferenceType        *string
	SortIndex            *int64
	Source               *string
}

func (p IndustryDataReferenceDefinitionOperationPredicate) Matches(input IndustryDataReferenceDefinition) bool {

	if p.Code != nil && (input.Code == nil || *p.Code != *input.Code) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsDisabled != nil && (input.IsDisabled == nil || *p.IsDisabled != *input.IsDisabled) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ReferenceType != nil && (input.ReferenceType == nil || *p.ReferenceType != *input.ReferenceType) {
		return false
	}

	if p.SortIndex != nil && (input.SortIndex == nil || *p.SortIndex != *input.SortIndex) {
		return false
	}

	if p.Source != nil && (input.Source == nil || *p.Source != *input.Source) {
		return false
	}

	return true
}
