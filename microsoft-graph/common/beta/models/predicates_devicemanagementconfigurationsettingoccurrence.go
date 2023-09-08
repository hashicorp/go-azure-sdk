package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementConfigurationSettingOccurrenceOperationPredicate struct {
	MaxDeviceOccurrence *int64
	MinDeviceOccurrence *int64
	ODataType           *string
}

func (p DeviceManagementConfigurationSettingOccurrenceOperationPredicate) Matches(input DeviceManagementConfigurationSettingOccurrence) bool {

	if p.MaxDeviceOccurrence != nil && (input.MaxDeviceOccurrence == nil || *p.MaxDeviceOccurrence != *input.MaxDeviceOccurrence) {
		return false
	}

	if p.MinDeviceOccurrence != nil && (input.MinDeviceOccurrence == nil || *p.MinDeviceOccurrence != *input.MinDeviceOccurrence) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
