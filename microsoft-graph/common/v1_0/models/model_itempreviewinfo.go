package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ItemPreviewInfo struct {
	GetUrl         *string `json:"getUrl,omitempty"`
	ODataType      *string `json:"@odata.type,omitempty"`
	PostParameters *string `json:"postParameters,omitempty"`
	PostUrl        *string `json:"postUrl,omitempty"`
}
