package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyPresentationMultiTextBoxOperationPredicate struct {
	Id                   *string
	Label                *string
	LastModifiedDateTime *string
	MaxLength            *int64
	MaxStrings           *int64
	ODataType            *string
	Required             *bool
}

func (p GroupPolicyPresentationMultiTextBoxOperationPredicate) Matches(input GroupPolicyPresentationMultiTextBox) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Label != nil && (input.Label == nil || *p.Label != *input.Label) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.MaxLength != nil && (input.MaxLength == nil || *p.MaxLength != *input.MaxLength) {
		return false
	}

	if p.MaxStrings != nil && (input.MaxStrings == nil || *p.MaxStrings != *input.MaxStrings) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Required != nil && (input.Required == nil || *p.Required != *input.Required) {
		return false
	}

	return true
}
