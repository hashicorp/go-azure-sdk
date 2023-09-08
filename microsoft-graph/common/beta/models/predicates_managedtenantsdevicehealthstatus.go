package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedTenantsDeviceHealthStatusOperationPredicate struct {
	BlueScreenCount     *int64
	DeviceId            *string
	DeviceMake          *string
	DeviceModel         *string
	DeviceName          *string
	HealthStatus        *string
	Id                  *string
	LastUpdatedDateTime *string
	ODataType           *string
	OsVersion           *string
	PrimaryDiskType     *string
	RestartCount        *int64
	TenantDisplayName   *string
	TenantId            *string
	TopProcesses        *string
}

func (p ManagedTenantsDeviceHealthStatusOperationPredicate) Matches(input ManagedTenantsDeviceHealthStatus) bool {

	if p.BlueScreenCount != nil && (input.BlueScreenCount == nil || *p.BlueScreenCount != *input.BlueScreenCount) {
		return false
	}

	if p.DeviceId != nil && (input.DeviceId == nil || *p.DeviceId != *input.DeviceId) {
		return false
	}

	if p.DeviceMake != nil && (input.DeviceMake == nil || *p.DeviceMake != *input.DeviceMake) {
		return false
	}

	if p.DeviceModel != nil && (input.DeviceModel == nil || *p.DeviceModel != *input.DeviceModel) {
		return false
	}

	if p.DeviceName != nil && (input.DeviceName == nil || *p.DeviceName != *input.DeviceName) {
		return false
	}

	if p.HealthStatus != nil && (input.HealthStatus == nil || *p.HealthStatus != *input.HealthStatus) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.LastUpdatedDateTime != nil && (input.LastUpdatedDateTime == nil || *p.LastUpdatedDateTime != *input.LastUpdatedDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OsVersion != nil && (input.OsVersion == nil || *p.OsVersion != *input.OsVersion) {
		return false
	}

	if p.PrimaryDiskType != nil && (input.PrimaryDiskType == nil || *p.PrimaryDiskType != *input.PrimaryDiskType) {
		return false
	}

	if p.RestartCount != nil && (input.RestartCount == nil || *p.RestartCount != *input.RestartCount) {
		return false
	}

	if p.TenantDisplayName != nil && (input.TenantDisplayName == nil || *p.TenantDisplayName != *input.TenantDisplayName) {
		return false
	}

	if p.TenantId != nil && (input.TenantId == nil || *p.TenantId != *input.TenantId) {
		return false
	}

	if p.TopProcesses != nil && (input.TopProcesses == nil || *p.TopProcesses != *input.TopProcesses) {
		return false
	}

	return true
}
