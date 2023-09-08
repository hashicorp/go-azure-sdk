package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookWorksheet struct {
	Charts      *[]WorkbookChart             `json:"charts,omitempty"`
	Id          *string                      `json:"id,omitempty"`
	Name        *string                      `json:"name,omitempty"`
	Names       *[]WorkbookNamedItem         `json:"names,omitempty"`
	ODataType   *string                      `json:"@odata.type,omitempty"`
	PivotTables *[]WorkbookPivotTable        `json:"pivotTables,omitempty"`
	Position    *int64                       `json:"position,omitempty"`
	Protection  *WorkbookWorksheetProtection `json:"protection,omitempty"`
	Tables      *[]WorkbookTable             `json:"tables,omitempty"`
	Visibility  *string                      `json:"visibility,omitempty"`
}
