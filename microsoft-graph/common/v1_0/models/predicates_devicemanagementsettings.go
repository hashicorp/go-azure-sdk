package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementSettingsOperationPredicate struct {
	DeviceComplianceCheckinThresholdDays *int64
	IsScheduledActionEnabled             *bool
	ODataType                            *string
	SecureByDefault                      *bool
}

func (p DeviceManagementSettingsOperationPredicate) Matches(input DeviceManagementSettings) bool {

	if p.DeviceComplianceCheckinThresholdDays != nil && (input.DeviceComplianceCheckinThresholdDays == nil || *p.DeviceComplianceCheckinThresholdDays != *input.DeviceComplianceCheckinThresholdDays) {
		return false
	}

	if p.IsScheduledActionEnabled != nil && (input.IsScheduledActionEnabled == nil || *p.IsScheduledActionEnabled != *input.IsScheduledActionEnabled) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.SecureByDefault != nil && (input.SecureByDefault == nil || *p.SecureByDefault != *input.SecureByDefault) {
		return false
	}

	return true
}
