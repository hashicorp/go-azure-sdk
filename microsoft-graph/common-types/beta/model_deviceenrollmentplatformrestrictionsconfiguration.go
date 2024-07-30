package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceEnrollmentPlatformRestrictionsConfiguration struct {
	AndroidForWorkRestriction         *DeviceEnrollmentPlatformRestriction                                                `json:"androidForWorkRestriction,omitempty"`
	AndroidRestriction                *DeviceEnrollmentPlatformRestriction                                                `json:"androidRestriction,omitempty"`
	Assignments                       *[]EnrollmentConfigurationAssignment                                                `json:"assignments,omitempty"`
	CreatedDateTime                   *string                                                                             `json:"createdDateTime,omitempty"`
	Description                       *string                                                                             `json:"description,omitempty"`
	DeviceEnrollmentConfigurationType *DeviceEnrollmentPlatformRestrictionsConfigurationDeviceEnrollmentConfigurationType `json:"deviceEnrollmentConfigurationType,omitempty"`
	DisplayName                       *string                                                                             `json:"displayName,omitempty"`
	Id                                *string                                                                             `json:"id,omitempty"`
	IosRestriction                    *DeviceEnrollmentPlatformRestriction                                                `json:"iosRestriction,omitempty"`
	LastModifiedDateTime              *string                                                                             `json:"lastModifiedDateTime,omitempty"`
	MacOSRestriction                  *DeviceEnrollmentPlatformRestriction                                                `json:"macOSRestriction,omitempty"`
	MacRestriction                    *DeviceEnrollmentPlatformRestriction                                                `json:"macRestriction,omitempty"`
	ODataType                         *string                                                                             `json:"@odata.type,omitempty"`
	Priority                          *int64                                                                              `json:"priority,omitempty"`
	RoleScopeTagIds                   *[]string                                                                           `json:"roleScopeTagIds,omitempty"`
	Version                           *int64                                                                              `json:"version,omitempty"`
	WindowsHomeSkuRestriction         *DeviceEnrollmentPlatformRestriction                                                `json:"windowsHomeSkuRestriction,omitempty"`
	WindowsMobileRestriction          *DeviceEnrollmentPlatformRestriction                                                `json:"windowsMobileRestriction,omitempty"`
	WindowsRestriction                *DeviceEnrollmentPlatformRestriction                                                `json:"windowsRestriction,omitempty"`
}
