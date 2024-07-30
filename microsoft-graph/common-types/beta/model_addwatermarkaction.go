package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AddWatermarkAction struct {
	FontColor     *string                   `json:"fontColor,omitempty"`
	FontName      *string                   `json:"fontName,omitempty"`
	FontSize      *int64                    `json:"fontSize,omitempty"`
	Layout        *AddWatermarkActionLayout `json:"layout,omitempty"`
	ODataType     *string                   `json:"@odata.type,omitempty"`
	Text          *string                   `json:"text,omitempty"`
	UiElementName *string                   `json:"uiElementName,omitempty"`
}
