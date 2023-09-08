package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SecurityOcrSettings struct {
	IsEnabled    *bool   `json:"isEnabled,omitempty"`
	MaxImageSize *int64  `json:"maxImageSize,omitempty"`
	ODataType    *string `json:"@odata.type,omitempty"`
	Timeout      *string `json:"timeout,omitempty"`
}
