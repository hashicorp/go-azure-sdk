package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkbookRangeBorder struct {
	Color     *string `json:"color,omitempty"`
	Id        *string `json:"id,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	SideIndex *string `json:"sideIndex,omitempty"`
	Style     *string `json:"style,omitempty"`
	Weight    *string `json:"weight,omitempty"`
}
