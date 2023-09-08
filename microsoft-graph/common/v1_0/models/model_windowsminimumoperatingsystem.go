package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsMinimumOperatingSystem struct {
	ODataType *string `json:"@odata.type,omitempty"`
	V100      *bool   `json:"v10_0,omitempty"`
	V80       *bool   `json:"v8_0,omitempty"`
	V81       *bool   `json:"v8_1,omitempty"`
}
