package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookRangeFormat struct {
	Borders             *[]WorkbookRangeBorder    `json:"borders,omitempty"`
	ColumnWidth         *float64                  `json:"columnWidth,omitempty"`
	Fill                *WorkbookRangeFill        `json:"fill,omitempty"`
	Font                *WorkbookRangeFont        `json:"font,omitempty"`
	HorizontalAlignment *string                   `json:"horizontalAlignment,omitempty"`
	Id                  *string                   `json:"id,omitempty"`
	ODataType           *string                   `json:"@odata.type,omitempty"`
	Protection          *WorkbookFormatProtection `json:"protection,omitempty"`
	RowHeight           *float64                  `json:"rowHeight,omitempty"`
	VerticalAlignment   *string                   `json:"verticalAlignment,omitempty"`
	WrapText            *bool                     `json:"wrapText,omitempty"`
}
