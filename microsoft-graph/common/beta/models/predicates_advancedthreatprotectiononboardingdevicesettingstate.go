package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AdvancedThreatProtectionOnboardingDeviceSettingStateOperationPredicate struct {
	ComplianceGracePeriodExpirationDateTime *string
	DeviceId                                *string
	DeviceModel                             *string
	DeviceName                              *string
	Id                                      *string
	ODataType                               *string
	Setting                                 *string
	SettingName                             *string
	UserEmail                               *string
	UserId                                  *string
	UserName                                *string
	UserPrincipalName                       *string
}

func (p AdvancedThreatProtectionOnboardingDeviceSettingStateOperationPredicate) Matches(input AdvancedThreatProtectionOnboardingDeviceSettingState) bool {

	if p.ComplianceGracePeriodExpirationDateTime != nil && (input.ComplianceGracePeriodExpirationDateTime == nil || *p.ComplianceGracePeriodExpirationDateTime != *input.ComplianceGracePeriodExpirationDateTime) {
		return false
	}

	if p.DeviceId != nil && (input.DeviceId == nil || *p.DeviceId != *input.DeviceId) {
		return false
	}

	if p.DeviceModel != nil && (input.DeviceModel == nil || *p.DeviceModel != *input.DeviceModel) {
		return false
	}

	if p.DeviceName != nil && (input.DeviceName == nil || *p.DeviceName != *input.DeviceName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.Setting != nil && (input.Setting == nil || *p.Setting != *input.Setting) {
		return false
	}

	if p.SettingName != nil && (input.SettingName == nil || *p.SettingName != *input.SettingName) {
		return false
	}

	if p.UserEmail != nil && (input.UserEmail == nil || *p.UserEmail != *input.UserEmail) {
		return false
	}

	if p.UserId != nil && (input.UserId == nil || *p.UserId != *input.UserId) {
		return false
	}

	if p.UserName != nil && (input.UserName == nil || *p.UserName != *input.UserName) {
		return false
	}

	if p.UserPrincipalName != nil && (input.UserPrincipalName == nil || *p.UserPrincipalName != *input.UserPrincipalName) {
		return false
	}

	return true
}
