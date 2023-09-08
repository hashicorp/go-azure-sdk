package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type UserExperienceAnalyticsWorkFromAnywhereDevicesSummaryOperationPredicate struct {
	CoManagedDevices                       *int64
	DevicesNotAutopilotRegistered          *int64
	DevicesWithoutAutopilotProfileAssigned *int64
	DevicesWithoutCloudIdentity            *int64
	IntuneDevices                          *int64
	ODataType                              *string
	TenantAttachDevices                    *int64
	TotalDevices                           *int64
	UnsupportedOSversionDevices            *int64
	Windows10Devices                       *int64
	Windows10DevicesWithoutTenantAttach    *int64
}

func (p UserExperienceAnalyticsWorkFromAnywhereDevicesSummaryOperationPredicate) Matches(input UserExperienceAnalyticsWorkFromAnywhereDevicesSummary) bool {

	if p.CoManagedDevices != nil && (input.CoManagedDevices == nil || *p.CoManagedDevices != *input.CoManagedDevices) {
		return false
	}

	if p.DevicesNotAutopilotRegistered != nil && (input.DevicesNotAutopilotRegistered == nil || *p.DevicesNotAutopilotRegistered != *input.DevicesNotAutopilotRegistered) {
		return false
	}

	if p.DevicesWithoutAutopilotProfileAssigned != nil && (input.DevicesWithoutAutopilotProfileAssigned == nil || *p.DevicesWithoutAutopilotProfileAssigned != *input.DevicesWithoutAutopilotProfileAssigned) {
		return false
	}

	if p.DevicesWithoutCloudIdentity != nil && (input.DevicesWithoutCloudIdentity == nil || *p.DevicesWithoutCloudIdentity != *input.DevicesWithoutCloudIdentity) {
		return false
	}

	if p.IntuneDevices != nil && (input.IntuneDevices == nil || *p.IntuneDevices != *input.IntuneDevices) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.TenantAttachDevices != nil && (input.TenantAttachDevices == nil || *p.TenantAttachDevices != *input.TenantAttachDevices) {
		return false
	}

	if p.TotalDevices != nil && (input.TotalDevices == nil || *p.TotalDevices != *input.TotalDevices) {
		return false
	}

	if p.UnsupportedOSversionDevices != nil && (input.UnsupportedOSversionDevices == nil || *p.UnsupportedOSversionDevices != *input.UnsupportedOSversionDevices) {
		return false
	}

	if p.Windows10Devices != nil && (input.Windows10Devices == nil || *p.Windows10Devices != *input.Windows10Devices) {
		return false
	}

	if p.Windows10DevicesWithoutTenantAttach != nil && (input.Windows10DevicesWithoutTenantAttach == nil || *p.Windows10DevicesWithoutTenantAttach != *input.Windows10DevicesWithoutTenantAttach) {
		return false
	}

	return true
}
