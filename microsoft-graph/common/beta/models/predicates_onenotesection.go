package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OnenoteSectionOperationPredicate struct {
	CreatedDateTime      *string
	DisplayName          *string
	Id                   *string
	IsDefault            *bool
	LastModifiedDateTime *string
	ODataType            *string
	PagesUrl             *string
	Self                 *string
}

func (p OnenoteSectionOperationPredicate) Matches(input OnenoteSection) bool {

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DisplayName != nil && (input.DisplayName == nil || *p.DisplayName != *input.DisplayName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.IsDefault != nil && (input.IsDefault == nil || *p.IsDefault != *input.IsDefault) {
		return false
	}

	if p.LastModifiedDateTime != nil && (input.LastModifiedDateTime == nil || *p.LastModifiedDateTime != *input.LastModifiedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PagesUrl != nil && (input.PagesUrl == nil || *p.PagesUrl != *input.PagesUrl) {
		return false
	}

	if p.Self != nil && (input.Self == nil || *p.Self != *input.Self) {
		return false
	}

	return true
}
