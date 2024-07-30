package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookRange struct {
	Address       *string              `json:"address,omitempty"`
	AddressLocal  *string              `json:"addressLocal,omitempty"`
	CellCount     *int64               `json:"cellCount,omitempty"`
	ColumnCount   *int64               `json:"columnCount,omitempty"`
	ColumnHidden  *bool                `json:"columnHidden,omitempty"`
	ColumnIndex   *int64               `json:"columnIndex,omitempty"`
	Format        *WorkbookRangeFormat `json:"format,omitempty"`
	Formulas      *Json                `json:"formulas,omitempty"`
	FormulasLocal *Json                `json:"formulasLocal,omitempty"`
	FormulasR1C1  *Json                `json:"formulasR1C1,omitempty"`
	Hidden        *bool                `json:"hidden,omitempty"`
	Id            *string              `json:"id,omitempty"`
	NumberFormat  *Json                `json:"numberFormat,omitempty"`
	ODataType     *string              `json:"@odata.type,omitempty"`
	RowCount      *int64               `json:"rowCount,omitempty"`
	RowHidden     *bool                `json:"rowHidden,omitempty"`
	RowIndex      *int64               `json:"rowIndex,omitempty"`
	Sort          *WorkbookRangeSort   `json:"sort,omitempty"`
	Text          *Json                `json:"text,omitempty"`
	ValueTypes    *Json                `json:"valueTypes,omitempty"`
	Values        *Json                `json:"values,omitempty"`
	Worksheet     *WorkbookWorksheet   `json:"worksheet,omitempty"`
}
