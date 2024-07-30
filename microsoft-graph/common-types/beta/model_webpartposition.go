package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WebPartPosition struct {
	ColumnId            *float64 `json:"columnId,omitempty"`
	HorizontalSectionId *float64 `json:"horizontalSectionId,omitempty"`
	IsInVerticalSection *bool    `json:"isInVerticalSection,omitempty"`
	ODataType           *string  `json:"@odata.type,omitempty"`
	WebPartIndex        *float64 `json:"webPartIndex,omitempty"`
}
