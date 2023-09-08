package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VisualInfo struct {
	Attribution     *ImageInfo `json:"attribution,omitempty"`
	BackgroundColor *string    `json:"backgroundColor,omitempty"`
	Content         *Json      `json:"content,omitempty"`
	Description     *string    `json:"description,omitempty"`
	DisplayText     *string    `json:"displayText,omitempty"`
	ODataType       *string    `json:"@odata.type,omitempty"`
}
