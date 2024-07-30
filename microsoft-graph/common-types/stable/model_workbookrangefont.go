package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookRangeFont struct {
	Bold      *bool    `json:"bold,omitempty"`
	Color     *string  `json:"color,omitempty"`
	Id        *string  `json:"id,omitempty"`
	Italic    *bool    `json:"italic,omitempty"`
	Name      *string  `json:"name,omitempty"`
	ODataType *string  `json:"@odata.type,omitempty"`
	Size      *float64 `json:"size,omitempty"`
	Underline *string  `json:"underline,omitempty"`
}
