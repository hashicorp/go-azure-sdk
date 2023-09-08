package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceManagementSettingsOperationPredicate struct {
	AndroidDeviceAdministratorEnrollmentEnabled *bool
	DerivedCredentialUrl                        *string
	DeviceComplianceCheckinThresholdDays        *int64
	DeviceInactivityBeforeRetirementInDay       *int64
	EnableAutopilotDiagnostics                  *bool
	EnableDeviceGroupMembershipReport           *bool
	EnableEnhancedTroubleshootingExperience     *bool
	EnableLogCollection                         *bool
	EnhancedJailBreak                           *bool
	IgnoreDevicesForUnsupportedSettingsEnabled  *bool
	IsScheduledActionEnabled                    *bool
	ODataType                                   *string
	SecureByDefault                             *bool
}

func (p DeviceManagementSettingsOperationPredicate) Matches(input DeviceManagementSettings) bool {

	if p.AndroidDeviceAdministratorEnrollmentEnabled != nil && (input.AndroidDeviceAdministratorEnrollmentEnabled == nil || *p.AndroidDeviceAdministratorEnrollmentEnabled != *input.AndroidDeviceAdministratorEnrollmentEnabled) {
		return false
	}

	if p.DerivedCredentialUrl != nil && (input.DerivedCredentialUrl == nil || *p.DerivedCredentialUrl != *input.DerivedCredentialUrl) {
		return false
	}

	if p.DeviceComplianceCheckinThresholdDays != nil && (input.DeviceComplianceCheckinThresholdDays == nil || *p.DeviceComplianceCheckinThresholdDays != *input.DeviceComplianceCheckinThresholdDays) {
		return false
	}

	if p.DeviceInactivityBeforeRetirementInDay != nil && (input.DeviceInactivityBeforeRetirementInDay == nil || *p.DeviceInactivityBeforeRetirementInDay != *input.DeviceInactivityBeforeRetirementInDay) {
		return false
	}

	if p.EnableAutopilotDiagnostics != nil && (input.EnableAutopilotDiagnostics == nil || *p.EnableAutopilotDiagnostics != *input.EnableAutopilotDiagnostics) {
		return false
	}

	if p.EnableDeviceGroupMembershipReport != nil && (input.EnableDeviceGroupMembershipReport == nil || *p.EnableDeviceGroupMembershipReport != *input.EnableDeviceGroupMembershipReport) {
		return false
	}

	if p.EnableEnhancedTroubleshootingExperience != nil && (input.EnableEnhancedTroubleshootingExperience == nil || *p.EnableEnhancedTroubleshootingExperience != *input.EnableEnhancedTroubleshootingExperience) {
		return false
	}

	if p.EnableLogCollection != nil && (input.EnableLogCollection == nil || *p.EnableLogCollection != *input.EnableLogCollection) {
		return false
	}

	if p.EnhancedJailBreak != nil && (input.EnhancedJailBreak == nil || *p.EnhancedJailBreak != *input.EnhancedJailBreak) {
		return false
	}

	if p.IgnoreDevicesForUnsupportedSettingsEnabled != nil && (input.IgnoreDevicesForUnsupportedSettingsEnabled == nil || *p.IgnoreDevicesForUnsupportedSettingsEnabled != *input.IgnoreDevicesForUnsupportedSettingsEnabled) {
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
