package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CanvasLayout struct {
	HorizontalSections *[]HorizontalSection `json:"horizontalSections,omitempty"`
	Id                 *string              `json:"id,omitempty"`
	ODataType          *string              `json:"@odata.type,omitempty"`
	VerticalSection    *VerticalSection     `json:"verticalSection,omitempty"`
}
