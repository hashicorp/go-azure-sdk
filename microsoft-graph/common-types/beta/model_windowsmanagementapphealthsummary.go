package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagementAppHealthSummary struct {
	HealthyDeviceCount   *int64  `json:"healthyDeviceCount,omitempty"`
	Id                   *string `json:"id,omitempty"`
	ODataType            *string `json:"@odata.type,omitempty"`
	UnhealthyDeviceCount *int64  `json:"unhealthyDeviceCount,omitempty"`
	UnknownDeviceCount   *int64  `json:"unknownDeviceCount,omitempty"`
}
