package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSMinimumOperatingSystem struct {
	ODataType *string `json:"@odata.type,omitempty"`
	V1010     *bool   `json:"v10_10,omitempty"`
	V1011     *bool   `json:"v10_11,omitempty"`
	V1012     *bool   `json:"v10_12,omitempty"`
	V1013     *bool   `json:"v10_13,omitempty"`
	V1014     *bool   `json:"v10_14,omitempty"`
	V1015     *bool   `json:"v10_15,omitempty"`
	V107      *bool   `json:"v10_7,omitempty"`
	V108      *bool   `json:"v10_8,omitempty"`
	V109      *bool   `json:"v10_9,omitempty"`
	V110      *bool   `json:"v11_0,omitempty"`
	V120      *bool   `json:"v12_0,omitempty"`
	V130      *bool   `json:"v13_0,omitempty"`
}
