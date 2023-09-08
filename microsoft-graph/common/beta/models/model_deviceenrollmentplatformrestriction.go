package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentPlatformRestriction struct {
	BlockedManufacturers            *[]string `json:"blockedManufacturers,omitempty"`
	BlockedSkus                     *[]string `json:"blockedSkus,omitempty"`
	ODataType                       *string   `json:"@odata.type,omitempty"`
	OsMaximumVersion                *string   `json:"osMaximumVersion,omitempty"`
	OsMinimumVersion                *string   `json:"osMinimumVersion,omitempty"`
	PersonalDeviceEnrollmentBlocked *bool     `json:"personalDeviceEnrollmentBlocked,omitempty"`
	PlatformBlocked                 *bool     `json:"platformBlocked,omitempty"`
}
