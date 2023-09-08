package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupPolicyPresentationListBoxOperationPredicate struct {
	ExplicitValue        *bool
	Id                   *string
	Label                *string
	LastModifiedDateTime *string
	ODataType            *string
	ValuePrefix          *string
}

func (p GroupPolicyPresentationListBoxOperationPredicate) Matches(input GroupPolicyPresentationListBox) bool {

	if p.ExplicitValue != nil && (input.ExplicitValue == nil || *p.ExplicitValue != *input.ExplicitValue) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.Label != nil && (input.Label == nil || *p.Label != *input.Label) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ValuePrefix != nil && (input.ValuePrefix == nil || *p.ValuePrefix != *input.ValuePrefix) {
		return false
	}

	return true
}
