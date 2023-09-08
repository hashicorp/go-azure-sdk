package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookChartTitleOperationPredicate struct {
	Id        *string
	ODataType *string
	Overlay   *bool
	Text      *string
	Visible   *bool
}

func (p WorkbookChartTitleOperationPredicate) Matches(input WorkbookChartTitle) bool {

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Overlay != nil && (input.Overlay == nil || *p.Overlay != *input.Overlay) {
		return false
	}

	if p.Text != nil && (input.Text == nil || *p.Text != *input.Text) {
		return false
	}

	if p.Visible != nil && (input.Visible == nil || *p.Visible != *input.Visible) {
		return false
	}

	return true
}
