package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChromeOSDeviceProperty struct {
	Name      *string `json:"name,omitempty"`
	ODataType *string `json:"@odata.type,omitempty"`
	Updatable *bool   `json:"updatable,omitempty"`
	Value     *string `json:"value,omitempty"`
	ValueType *string `json:"valueType,omitempty"`
}
