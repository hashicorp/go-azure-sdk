package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDeviceMobileAppConfiguration struct {
	Assignments          *[]ManagedDeviceMobileAppConfigurationAssignment   `json:"assignments,omitempty"`
	CreatedDateTime      *string                                            `json:"createdDateTime,omitempty"`
	Description          *string                                            `json:"description,omitempty"`
	DeviceStatusSummary  *ManagedDeviceMobileAppConfigurationDeviceSummary  `json:"deviceStatusSummary,omitempty"`
	DeviceStatuses       *[]ManagedDeviceMobileAppConfigurationDeviceStatus `json:"deviceStatuses,omitempty"`
	DisplayName          *string                                            `json:"displayName,omitempty"`
	Id                   *string                                            `json:"id,omitempty"`
	LastModifiedDateTime *string                                            `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                                            `json:"@odata.type,omitempty"`
	TargetedMobileApps   *[]string                                          `json:"targetedMobileApps,omitempty"`
	UserStatusSummary    *ManagedDeviceMobileAppConfigurationUserSummary    `json:"userStatusSummary,omitempty"`
	UserStatuses         *[]ManagedDeviceMobileAppConfigurationUserStatus   `json:"userStatuses,omitempty"`
	Version              *int64                                             `json:"version,omitempty"`
}
