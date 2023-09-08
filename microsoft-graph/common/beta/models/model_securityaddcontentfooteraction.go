package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityAddContentFooterAction struct {
	Alignment     *SecurityAddContentFooterActionAlignment `json:"alignment,omitempty"`
	FontColor     *string                                  `json:"fontColor,omitempty"`
	FontName      *string                                  `json:"fontName,omitempty"`
	FontSize      *int64                                   `json:"fontSize,omitempty"`
	Margin        *int64                                   `json:"margin,omitempty"`
	ODataType     *string                                  `json:"@odata.type,omitempty"`
	Text          *string                                  `json:"text,omitempty"`
	UiElementName *string                                  `json:"uiElementName,omitempty"`
}
