package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookChartLegendOperationPredicate struct {
	Id        *string
	ODataType *string
	Overlay   *bool
	Position  *string
	Visible   *bool
}

func (p WorkbookChartLegendOperationPredicate) Matches(input WorkbookChartLegend) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Overlay != nil && (input.Overlay == nil || *p.Overlay != *input.Overlay) {
		return false
	}

	if p.Position != nil && (input.Position == nil || *p.Position != *input.Position) {
		return false
	}

	if p.Visible != nil && (input.Visible == nil || *p.Visible != *input.Visible) {
		return false
	}

	return true
}
