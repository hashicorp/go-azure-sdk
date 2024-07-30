package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkaccessDeviceUsageSummary struct {
	ActiveDeviceCount   *int64  `json:"activeDeviceCount,omitempty"`
	InactiveDeviceCount *int64  `json:"inactiveDeviceCount,omitempty"`
	ODataType           *string `json:"@odata.type,omitempty"`
	TotalDeviceCount    *int64  `json:"totalDeviceCount,omitempty"`
}
