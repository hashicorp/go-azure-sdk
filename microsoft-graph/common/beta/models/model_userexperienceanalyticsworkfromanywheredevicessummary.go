package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsWorkFromAnywhereDevicesSummary struct {
	AutopilotDevicesSummary                *UserExperienceAnalyticsAutopilotDevicesSummary       `json:"autopilotDevicesSummary,omitempty"`
	CloudIdentityDevicesSummary            *UserExperienceAnalyticsCloudIdentityDevicesSummary   `json:"cloudIdentityDevicesSummary,omitempty"`
	CloudManagementDevicesSummary          *UserExperienceAnalyticsCloudManagementDevicesSummary `json:"cloudManagementDevicesSummary,omitempty"`
	CoManagedDevices                       *int64                                                `json:"coManagedDevices,omitempty"`
	DevicesNotAutopilotRegistered          *int64                                                `json:"devicesNotAutopilotRegistered,omitempty"`
	DevicesWithoutAutopilotProfileAssigned *int64                                                `json:"devicesWithoutAutopilotProfileAssigned,omitempty"`
	DevicesWithoutCloudIdentity            *int64                                                `json:"devicesWithoutCloudIdentity,omitempty"`
	IntuneDevices                          *int64                                                `json:"intuneDevices,omitempty"`
	ODataType                              *string                                               `json:"@odata.type,omitempty"`
	TenantAttachDevices                    *int64                                                `json:"tenantAttachDevices,omitempty"`
	TotalDevices                           *int64                                                `json:"totalDevices,omitempty"`
	UnsupportedOSversionDevices            *int64                                                `json:"unsupportedOSversionDevices,omitempty"`
	Windows10Devices                       *int64                                                `json:"windows10Devices,omitempty"`
	Windows10DevicesSummary                *UserExperienceAnalyticsWindows10DevicesSummary       `json:"windows10DevicesSummary,omitempty"`
	Windows10DevicesWithoutTenantAttach    *int64                                                `json:"windows10DevicesWithoutTenantAttach,omitempty"`
}
