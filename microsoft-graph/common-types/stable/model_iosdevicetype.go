package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosDeviceType struct {
	IPad          *bool   `json:"iPad,omitempty"`
	IPhoneAndIPod *bool   `json:"iPhoneAndIPod,omitempty"`
	ODataType     *string `json:"@odata.type,omitempty"`
}
