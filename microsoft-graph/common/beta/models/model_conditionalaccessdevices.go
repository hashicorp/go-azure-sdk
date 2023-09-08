package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConditionalAccessDevices struct {
	DeviceFilter        *ConditionalAccessFilter `json:"deviceFilter,omitempty"`
	ExcludeDeviceStates *[]string                `json:"excludeDeviceStates,omitempty"`
	ExcludeDevices      *[]string                `json:"excludeDevices,omitempty"`
	IncludeDeviceStates *[]string                `json:"includeDeviceStates,omitempty"`
	IncludeDevices      *[]string                `json:"includeDevices,omitempty"`
	ODataType           *string                  `json:"@odata.type,omitempty"`
}
