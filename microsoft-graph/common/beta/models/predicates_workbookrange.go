package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookRangeOperationPredicate struct {
	Address      *string
	AddressLocal *string
	CellCount    *int64
	ColumnCount  *int64
	ColumnHidden *bool
	ColumnIndex  *int64
	Hidden       *bool
	Id           *string
	ODataType    *string
	RowCount     *int64
	RowHidden    *bool
	RowIndex     *int64
}

func (p WorkbookRangeOperationPredicate) Matches(input WorkbookRange) bool {

	if p.Address != nil && (input.Address == nil || *p.Address != *input.Address) {
		return false
	}

	if p.AddressLocal != nil && (input.AddressLocal == nil || *p.AddressLocal != *input.AddressLocal) {
		return false
	}

	if p.CellCount != nil && (input.CellCount == nil || *p.CellCount != *input.CellCount) {
		return false
	}

	if p.ColumnCount != nil && (input.ColumnCount == nil || *p.ColumnCount != *input.ColumnCount) {
		return false
	}

	if p.ColumnHidden != nil && (input.ColumnHidden == nil || *p.ColumnHidden != *input.ColumnHidden) {
		return false
	}

	if p.ColumnIndex != nil && (input.ColumnIndex == nil || *p.ColumnIndex != *input.ColumnIndex) {
		return false
	}

	if p.Hidden != nil && (input.Hidden == nil || *p.Hidden != *input.Hidden) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.RowCount != nil && (input.RowCount == nil || *p.RowCount != *input.RowCount) {
		return false
	}

	if p.RowHidden != nil && (input.RowHidden == nil || *p.RowHidden != *input.RowHidden) {
		return false
	}

	if p.RowIndex != nil && (input.RowIndex == nil || *p.RowIndex != *input.RowIndex) {
		return false
	}

	return true
}
