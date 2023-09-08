package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosMobileAppConfiguration struct {
	Assignments          *[]ManagedDeviceMobileAppConfigurationAssignment   `json:"assignments,omitempty"`
	CreatedDateTime      *string                                            `json:"createdDateTime,omitempty"`
	Description          *string                                            `json:"description,omitempty"`
	DeviceStatusSummary  *ManagedDeviceMobileAppConfigurationDeviceSummary  `json:"deviceStatusSummary,omitempty"`
	DeviceStatuses       *[]ManagedDeviceMobileAppConfigurationDeviceStatus `json:"deviceStatuses,omitempty"`
	DisplayName          *string                                            `json:"displayName,omitempty"`
	EncodedSettingXml    *string                                            `json:"encodedSettingXml,omitempty"`
	Id                   *string                                            `json:"id,omitempty"`
	LastModifiedDateTime *string                                            `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                                            `json:"@odata.type,omitempty"`
	RoleScopeTagIds      *[]string                                          `json:"roleScopeTagIds,omitempty"`
	Settings             *[]AppConfigurationSettingItem                     `json:"settings,omitempty"`
	TargetedMobileApps   *[]string                                          `json:"targetedMobileApps,omitempty"`
	UserStatusSummary    *ManagedDeviceMobileAppConfigurationUserSummary    `json:"userStatusSummary,omitempty"`
	UserStatuses         *[]ManagedDeviceMobileAppConfigurationUserStatus   `json:"userStatuses,omitempty"`
	Version              *int64                                             `json:"version,omitempty"`
}
