package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookRangeBorderOperationPredicate struct {
	Color     *string
	Id        *string
	ODataType *string
	SideIndex *string
	Style     *string
	Weight    *string
}

func (p WorkbookRangeBorderOperationPredicate) Matches(input WorkbookRangeBorder) bool {

	if p.Color != nil && (input.Color == nil || *p.Color != *input.Color) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SideIndex != nil && (input.SideIndex == nil || *p.SideIndex != *input.SideIndex) {
		return false
	}

	if p.Style != nil && (input.Style == nil || *p.Style != *input.Style) {
		return false
	}

	if p.Weight != nil && (input.Weight == nil || *p.Weight != *input.Weight) {
		return false
	}

	return true
}
