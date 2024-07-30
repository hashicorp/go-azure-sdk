package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookRangeView struct {
	CellAddresses *Json                `json:"cellAddresses,omitempty"`
	ColumnCount   *int64               `json:"columnCount,omitempty"`
	Formulas      *Json                `json:"formulas,omitempty"`
	FormulasLocal *Json                `json:"formulasLocal,omitempty"`
	FormulasR1C1  *Json                `json:"formulasR1C1,omitempty"`
	Id            *string              `json:"id,omitempty"`
	Index         *int64               `json:"index,omitempty"`
	NumberFormat  *Json                `json:"numberFormat,omitempty"`
	ODataType     *string              `json:"@odata.type,omitempty"`
	RowCount      *int64               `json:"rowCount,omitempty"`
	Rows          *[]WorkbookRangeView `json:"rows,omitempty"`
	Text          *Json                `json:"text,omitempty"`
	ValueTypes    *Json                `json:"valueTypes,omitempty"`
	Values        *Json                `json:"values,omitempty"`
}
