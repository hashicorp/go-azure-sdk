package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceCleanupSettingsOperationPredicate struct {
	DeviceInactivityBeforeRetirementInDays *string
	ODataType                              *string
}

func (p ManagedDeviceCleanupSettingsOperationPredicate) Matches(input ManagedDeviceCleanupSettings) bool {

	if p.DeviceInactivityBeforeRetirementInDays != nil && (input.DeviceInactivityBeforeRetirementInDays == nil || *p.DeviceInactivityBeforeRetirementInDays != *input.DeviceInactivityBeforeRetirementInDays) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
