package virtualendpointprovisioningpolicy

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookTable struct {
	Columns              *[]WorkbookTableColumn `json:"columns,omitempty"`
	HighlightFirstColumn *bool                  `json:"highlightFirstColumn,omitempty"`
	HighlightLastColumn  *bool                  `json:"highlightLastColumn,omitempty"`
	Id                   *string                `json:"id,omitempty"`
	LegacyId             *string                `json:"legacyId,omitempty"`
	Name                 *string                `json:"name,omitempty"`
	ODataType            *string                `json:"@odata.type,omitempty"`
	Rows                 *[]WorkbookTableRow    `json:"rows,omitempty"`
	ShowBandedColumns    *bool                  `json:"showBandedColumns,omitempty"`
	ShowBandedRows       *bool                  `json:"showBandedRows,omitempty"`
	ShowFilterButton     *bool                  `json:"showFilterButton,omitempty"`
	ShowHeaders          *bool                  `json:"showHeaders,omitempty"`
	ShowTotals           *bool                  `json:"showTotals,omitempty"`
	Sort                 *WorkbookTableSort     `json:"sort,omitempty"`
	Style                *string                `json:"style,omitempty"`
	Worksheet            *WorkbookWorksheet     `json:"worksheet,omitempty"`
}
