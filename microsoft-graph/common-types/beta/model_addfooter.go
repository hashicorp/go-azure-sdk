package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddFooter struct {
	Alignment *AddFooterAlignment `json:"alignment,omitempty"`
	FontColor *string             `json:"fontColor,omitempty"`
	FontSize  *int64              `json:"fontSize,omitempty"`
	Margin    *int64              `json:"margin,omitempty"`
	Name      *string             `json:"name,omitempty"`
	ODataType *string             `json:"@odata.type,omitempty"`
	Text      *string             `json:"text,omitempty"`
}
