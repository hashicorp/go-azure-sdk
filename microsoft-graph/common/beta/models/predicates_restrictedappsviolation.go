package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestrictedAppsViolationOperationPredicate struct {
	DeviceConfigurationId   *string
	DeviceConfigurationName *string
	DeviceName              *string
	Id                      *string
	ManagedDeviceId         *string
	ODataType               *string
	UserId                  *string
	UserName                *string
}

func (p RestrictedAppsViolationOperationPredicate) Matches(input RestrictedAppsViolation) bool {

	if p.DeviceConfigurationId != nil && (input.DeviceConfigurationId == nil || *p.DeviceConfigurationId != *input.DeviceConfigurationId) {
		return false
	}

	if p.DeviceConfigurationName != nil && (input.DeviceConfigurationName == nil || *p.DeviceConfigurationName != *input.DeviceConfigurationName) {
		return false
	}

	if p.DeviceName != nil && (input.DeviceName == nil || *p.DeviceName != *input.DeviceName) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.ManagedDeviceId != nil && (input.ManagedDeviceId == nil || *p.ManagedDeviceId != *input.ManagedDeviceId) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.UserId != nil && (input.UserId == nil || *p.UserId != *input.UserId) {
		return false
	}

	if p.UserName != nil && (input.UserName == nil || *p.UserName != *input.UserName) {
		return false
	}

	return true
}
