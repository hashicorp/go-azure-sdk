package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentPlatformRestrictionsConfiguration struct {
	AndroidRestriction       *DeviceEnrollmentPlatformRestriction `json:"androidRestriction,omitempty"`
	Assignments              *[]EnrollmentConfigurationAssignment `json:"assignments,omitempty"`
	CreatedDateTime          *string                              `json:"createdDateTime,omitempty"`
	Description              *string                              `json:"description,omitempty"`
	DisplayName              *string                              `json:"displayName,omitempty"`
	Id                       *string                              `json:"id,omitempty"`
	IosRestriction           *DeviceEnrollmentPlatformRestriction `json:"iosRestriction,omitempty"`
	LastModifiedDateTime     *string                              `json:"lastModifiedDateTime,omitempty"`
	MacOSRestriction         *DeviceEnrollmentPlatformRestriction `json:"macOSRestriction,omitempty"`
	ODataType                *string                              `json:"@odata.type,omitempty"`
	Priority                 *int64                               `json:"priority,omitempty"`
	Version                  *int64                               `json:"version,omitempty"`
	WindowsMobileRestriction *DeviceEnrollmentPlatformRestriction `json:"windowsMobileRestriction,omitempty"`
	WindowsRestriction       *DeviceEnrollmentPlatformRestriction `json:"windowsRestriction,omitempty"`
}
