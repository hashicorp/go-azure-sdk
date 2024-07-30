package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceExchangeAccessStateSummary struct {
	AllowedDeviceCount     *int64  `json:"allowedDeviceCount,omitempty"`
	BlockedDeviceCount     *int64  `json:"blockedDeviceCount,omitempty"`
	ODataType              *string `json:"@odata.type,omitempty"`
	QuarantinedDeviceCount *int64  `json:"quarantinedDeviceCount,omitempty"`
	UnavailableDeviceCount *int64  `json:"unavailableDeviceCount,omitempty"`
	UnknownDeviceCount     *int64  `json:"unknownDeviceCount,omitempty"`
}
