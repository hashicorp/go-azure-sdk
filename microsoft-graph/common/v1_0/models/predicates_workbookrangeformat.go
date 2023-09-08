package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookRangeFormatOperationPredicate struct {
	HorizontalAlignment *string
	Id                  *string
	ODataType           *string
	VerticalAlignment   *string
	WrapText            *bool
}

func (p WorkbookRangeFormatOperationPredicate) Matches(input WorkbookRangeFormat) bool {

	if p.HorizontalAlignment != nil && (input.HorizontalAlignment == nil || *p.HorizontalAlignment != *input.HorizontalAlignment) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.VerticalAlignment != nil && (input.VerticalAlignment == nil || *p.VerticalAlignment != *input.VerticalAlignment) {
		return false
	}

	if p.WrapText != nil && (input.WrapText == nil || *p.WrapText != *input.WrapText) {
		return false
	}

	return true
}
