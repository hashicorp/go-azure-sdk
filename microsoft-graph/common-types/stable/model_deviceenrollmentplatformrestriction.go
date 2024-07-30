package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentPlatformRestriction struct {
	ODataType                       *string `json:"@odata.type,omitempty"`
	OsMaximumVersion                *string `json:"osMaximumVersion,omitempty"`
	OsMinimumVersion                *string `json:"osMinimumVersion,omitempty"`
	PersonalDeviceEnrollmentBlocked *bool   `json:"personalDeviceEnrollmentBlocked,omitempty"`
	PlatformBlocked                 *bool   `json:"platformBlocked,omitempty"`
}
