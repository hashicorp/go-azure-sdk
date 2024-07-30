package site

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MobileAppSupportedDeviceType struct {
	MaximumOperatingSystemVersion *string                           `json:"maximumOperatingSystemVersion,omitempty"`
	MinimumOperatingSystemVersion *string                           `json:"minimumOperatingSystemVersion,omitempty"`
	ODataType                     *string                           `json:"@odata.type,omitempty"`
	Type                          *MobileAppSupportedDeviceTypeType `json:"type,omitempty"`
}
