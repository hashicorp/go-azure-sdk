package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsManagementAppHealthStateOperationPredicate struct {
	DeviceName          *string
	DeviceOSVersion     *string
	Id                  *string
	InstalledVersion    *string
	LastCheckInDateTime *string
	ODataType           *string
}

func (p WindowsManagementAppHealthStateOperationPredicate) Matches(input WindowsManagementAppHealthState) bool {

	if p.DeviceName != nil && (input.DeviceName == nil || *p.DeviceName != *input.DeviceName) {
		return false
	}

	if p.DeviceOSVersion != nil && (input.DeviceOSVersion == nil || *p.DeviceOSVersion != *input.DeviceOSVersion) {
		return false
	}

	if p.Id != nil && (input.Id == nil || *p.Id != *input.Id) {
		return false
	}

	if p.InstalledVersion != nil && (input.InstalledVersion == nil || *p.InstalledVersion != *input.InstalledVersion) {
		return false
	}

	if p.LastCheckInDateTime != nil && (input.LastCheckInDateTime == nil || *p.LastCheckInDateTime != *input.LastCheckInDateTime) {
		return false
	}

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	return true
}
