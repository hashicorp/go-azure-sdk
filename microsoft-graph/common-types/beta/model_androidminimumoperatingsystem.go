package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidMinimumOperatingSystem struct {
	ODataType *string `json:"@odata.type,omitempty"`
	V100      *bool   `json:"v10_0,omitempty"`
	V110      *bool   `json:"v11_0,omitempty"`
	V40       *bool   `json:"v4_0,omitempty"`
	V403      *bool   `json:"v4_0_3,omitempty"`
	V41       *bool   `json:"v4_1,omitempty"`
	V42       *bool   `json:"v4_2,omitempty"`
	V43       *bool   `json:"v4_3,omitempty"`
	V44       *bool   `json:"v4_4,omitempty"`
	V50       *bool   `json:"v5_0,omitempty"`
	V51       *bool   `json:"v5_1,omitempty"`
	V60       *bool   `json:"v6_0,omitempty"`
	V70       *bool   `json:"v7_0,omitempty"`
	V71       *bool   `json:"v7_1,omitempty"`
	V80       *bool   `json:"v8_0,omitempty"`
	V81       *bool   `json:"v8_1,omitempty"`
	V90       *bool   `json:"v9_0,omitempty"`
}
