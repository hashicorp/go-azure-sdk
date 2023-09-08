package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MarkContent struct {
	FontColor *string `json:"fontColor,omitempty"`
	FontSize  *int64  `json:"fontSize,omitempty"`
	Name      *string `json:"name,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	Text      *string `json:"text,omitempty"`
}
