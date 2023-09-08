package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookChartLineFormatOperationPredicate struct {
	Color     *string
	Id        *string
	ODataType *string
}

func (p WorkbookChartLineFormatOperationPredicate) Matches(input WorkbookChartLineFormat) bool {

	if p.Color != nil && (input.Color == nil || *p.Color != *input.Color) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
