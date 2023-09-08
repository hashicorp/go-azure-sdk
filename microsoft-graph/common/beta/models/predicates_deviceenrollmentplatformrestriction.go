package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentPlatformRestrictionOperationPredicate struct {
	ODataType                       *string
	OsMaximumVersion                *string
	OsMinimumVersion                *string
	PersonalDeviceEnrollmentBlocked *bool
	PlatformBlocked                 *bool
}

func (p DeviceEnrollmentPlatformRestrictionOperationPredicate) Matches(input DeviceEnrollmentPlatformRestriction) bool {

	if p.ODataType != nil && (input.ODataType == nil || *p.ODataType != *input.ODataType) {
		return false
	}

	if p.OsMaximumVersion != nil && (input.OsMaximumVersion == nil || *p.OsMaximumVersion != *input.OsMaximumVersion) {
		return false
	}

	if p.OsMinimumVersion != nil && (input.OsMinimumVersion == nil || *p.OsMinimumVersion != *input.OsMinimumVersion) {
		return false
	}

	if p.PersonalDeviceEnrollmentBlocked != nil && (input.PersonalDeviceEnrollmentBlocked == nil || *p.PersonalDeviceEnrollmentBlocked != *input.PersonalDeviceEnrollmentBlocked) {
		return false
	}

	if p.PlatformBlocked != nil && (input.PlatformBlocked == nil || *p.PlatformBlocked != *input.PlatformBlocked) {
		return false
	}

	return true
}
