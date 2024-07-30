package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsMinimumOperatingSystem struct {
	ODataType *string `json:"@odata.type,omitempty"`
	V100      *bool   `json:"v10_0,omitempty"`
	V101607   *bool   `json:"v10_1607,omitempty"`
	V101703   *bool   `json:"v10_1703,omitempty"`
	V101709   *bool   `json:"v10_1709,omitempty"`
	V101803   *bool   `json:"v10_1803,omitempty"`
	V101809   *bool   `json:"v10_1809,omitempty"`
	V101903   *bool   `json:"v10_1903,omitempty"`
	V101909   *bool   `json:"v10_1909,omitempty"`
	V102004   *bool   `json:"v10_2004,omitempty"`
	V1021H1   *bool   `json:"v10_21H1,omitempty"`
	V102H20   *bool   `json:"v10_2H20,omitempty"`
	V80       *bool   `json:"v8_0,omitempty"`
	V81       *bool   `json:"v8_1,omitempty"`
}
