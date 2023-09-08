package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HorizontalSection struct {
	Columns   *[]HorizontalSectionColumn `json:"columns,omitempty"`
	Emphasis  *HorizontalSectionEmphasis `json:"emphasis,omitempty"`
	Id        *string                    `json:"id,omitempty"`
	Layout    *HorizontalSectionLayout   `json:"layout,omitempty"`
	ODataType *string                    `json:"@odata.type,omitempty"`
}
