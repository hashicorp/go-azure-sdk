package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessDevice struct {
	DeviceId           *string                         `json:"deviceId,omitempty"`
	DisplayName        *string                         `json:"displayName,omitempty"`
	IsCompliant        *bool                           `json:"isCompliant,omitempty"`
	LastAccessDateTime *string                         `json:"lastAccessDateTime,omitempty"`
	ODataType          *string                         `json:"@odata.type,omitempty"`
	OperatingSystem    *string                         `json:"operatingSystem,omitempty"`
	TrafficType        *NetworkaccessDeviceTrafficType `json:"trafficType,omitempty"`
}
