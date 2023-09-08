package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosMinimumOperatingSystem struct {
	ODataType *string `json:"@odata.type,omitempty"`
	V100      *bool   `json:"v10_0,omitempty"`
	V110      *bool   `json:"v11_0,omitempty"`
	V120      *bool   `json:"v12_0,omitempty"`
	V130      *bool   `json:"v13_0,omitempty"`
	V140      *bool   `json:"v14_0,omitempty"`
	V150      *bool   `json:"v15_0,omitempty"`
	V80       *bool   `json:"v8_0,omitempty"`
	V90       *bool   `json:"v9_0,omitempty"`
}
