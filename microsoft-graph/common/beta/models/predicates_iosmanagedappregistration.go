package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosManagedAppRegistrationOperationPredicate struct {
	ApplicationVersion   *string
	AzureADDeviceId      *string
	CreatedDateTime      *string
	DeviceManufacturer   *string
	DeviceModel          *string
	DeviceName           *string
	DeviceTag            *string
	DeviceType           *string
	Id                   *string
	LastSyncDateTime     *string
	ManagedDeviceId      *string
	ManagementSdkVersion *string
	ODataType            *string
	PlatformVersion      *string
	UserId               *string
	Version              *string
}

func (p IosManagedAppRegistrationOperationPredicate) Matches(input IosManagedAppRegistration) bool {

	if p.ApplicationVersion != nil && (input.ApplicationVersion == nil || *p.ApplicationVersion != *input.ApplicationVersion) {
		return false
	}

	if p.AzureADDeviceId != nil && (input.AzureADDeviceId == nil || *p.AzureADDeviceId != *input.AzureADDeviceId) {
		return false
	}

	if p.CreatedDateTime != nil && (input.CreatedDateTime == nil || *p.CreatedDateTime != *input.CreatedDateTime) {
		return false
	}

	if p.DeviceManufacturer != nil && (input.DeviceManufacturer == nil || *p.DeviceManufacturer != *input.DeviceManufacturer) {
		return false
	}

	if p.DeviceModel != nil && (input.DeviceModel == nil || *p.DeviceModel != *input.DeviceModel) {
		return false
	}

	if p.DeviceName != nil && (input.DeviceName == nil || *p.DeviceName != *input.DeviceName) {
		return false
	}

	if p.DeviceTag != nil && (input.DeviceTag == nil || *p.DeviceTag != *input.DeviceTag) {
		return false
	}

	if p.DeviceType != nil && (input.DeviceType == nil || *p.DeviceType != *input.DeviceType) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastSyncDateTime != nil && (input.LastSyncDateTime == nil || *p.LastSyncDateTime != *input.LastSyncDateTime) {
		return false
	}

	if p.ManagedDeviceId != nil && (input.ManagedDeviceId == nil || *p.ManagedDeviceId != *input.ManagedDeviceId) {
		return false
	}

	if p.ManagementSdkVersion != nil && (input.ManagementSdkVersion == nil || *p.ManagementSdkVersion != *input.ManagementSdkVersion) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.PlatformVersion != nil && (input.PlatformVersion == nil || *p.PlatformVersion != *input.PlatformVersion) {
		return false
	}

	if p.UserId != nil && (input.UserId == nil || *p.UserId != *input.UserId) {
		return false
	}

	if p.Version != nil && (input.Version == nil || *p.Version != *input.Version) {
		return false
	}

	return true
}
