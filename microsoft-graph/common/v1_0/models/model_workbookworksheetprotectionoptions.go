package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookWorksheetProtectionOptions struct {
	AllowAutoFilter       *bool   `json:"allowAutoFilter,omitempty"`
	AllowDeleteColumns    *bool   `json:"allowDeleteColumns,omitempty"`
	AllowDeleteRows       *bool   `json:"allowDeleteRows,omitempty"`
	AllowFormatCells      *bool   `json:"allowFormatCells,omitempty"`
	AllowFormatColumns    *bool   `json:"allowFormatColumns,omitempty"`
	AllowFormatRows       *bool   `json:"allowFormatRows,omitempty"`
	AllowInsertColumns    *bool   `json:"allowInsertColumns,omitempty"`
	AllowInsertHyperlinks *bool   `json:"allowInsertHyperlinks,omitempty"`
	AllowInsertRows       *bool   `json:"allowInsertRows,omitempty"`
	AllowPivotTables      *bool   `json:"allowPivotTables,omitempty"`
	AllowSort             *bool   `json:"allowSort,omitempty"`
	ODataType             *string `json:"@odata.type,omitempty"`
}
