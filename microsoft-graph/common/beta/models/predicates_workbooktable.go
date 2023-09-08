package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookTableOperationPredicate struct {
	HighlightFirstColumn *bool
	HighlightLastColumn  *bool
	Id                   *string
	LegacyId             *string
	Name                 *string
	ODataType            *string
	ShowBandedColumns    *bool
	ShowBandedRows       *bool
	ShowFilterButton     *bool
	ShowHeaders          *bool
	ShowTotals           *bool
	Style                *string
}

func (p WorkbookTableOperationPredicate) Matches(input WorkbookTable) bool {

	if p.HighlightFirstColumn != nil && (input.HighlightFirstColumn == nil || *p.HighlightFirstColumn != *input.HighlightFirstColumn) {
		return false
	}

	if p.HighlightLastColumn != nil && (input.HighlightLastColumn == nil || *p.HighlightLastColumn != *input.HighlightLastColumn) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LegacyId != nil && (input.LegacyId == nil || *p.LegacyId != *input.LegacyId) {
		return false
	}

	if p.Name != nil && (input.Name == nil || *p.Name != *input.Name) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.ShowBandedColumns != nil && (input.ShowBandedColumns == nil || *p.ShowBandedColumns != *input.ShowBandedColumns) {
		return false
	}

	if p.ShowBandedRows != nil && (input.ShowBandedRows == nil || *p.ShowBandedRows != *input.ShowBandedRows) {
		return false
	}

	if p.ShowFilterButton != nil && (input.ShowFilterButton == nil || *p.ShowFilterButton != *input.ShowFilterButton) {
		return false
	}

	if p.ShowHeaders != nil && (input.ShowHeaders == nil || *p.ShowHeaders != *input.ShowHeaders) {
		return false
	}

	if p.ShowTotals != nil && (input.ShowTotals == nil || *p.ShowTotals != *input.ShowTotals) {
		return false
	}

	if p.Style != nil && (input.Style == nil || *p.Style != *input.Style) {
		return false
	}

	return true
}
