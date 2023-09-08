package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AndroidManagedStoreAppConfiguration struct {
	AppSupportsOemConfig *bool                                                    `json:"appSupportsOemConfig,omitempty"`
	Assignments          *[]ManagedDeviceMobileAppConfigurationAssignment         `json:"assignments,omitempty"`
	ConnectedAppsEnabled *bool                                                    `json:"connectedAppsEnabled,omitempty"`
	CreatedDateTime      *string                                                  `json:"createdDateTime,omitempty"`
	Description          *string                                                  `json:"description,omitempty"`
	DeviceStatusSummary  *ManagedDeviceMobileAppConfigurationDeviceSummary        `json:"deviceStatusSummary,omitempty"`
	DeviceStatuses       *[]ManagedDeviceMobileAppConfigurationDeviceStatus       `json:"deviceStatuses,omitempty"`
	DisplayName          *string                                                  `json:"displayName,omitempty"`
	Id                   *string                                                  `json:"id,omitempty"`
	LastModifiedDateTime *string                                                  `json:"lastModifiedDateTime,omitempty"`
	ODataType            *string                                                  `json:"@odata.type,omitempty"`
	PackageId            *string                                                  `json:"packageId,omitempty"`
	PayloadJson          *string                                                  `json:"payloadJson,omitempty"`
	PermissionActions    *[]AndroidPermissionAction                               `json:"permissionActions,omitempty"`
	ProfileApplicability *AndroidManagedStoreAppConfigurationProfileApplicability `json:"profileApplicability,omitempty"`
	RoleScopeTagIds      *[]string                                                `json:"roleScopeTagIds,omitempty"`
	TargetedMobileApps   *[]string                                                `json:"targetedMobileApps,omitempty"`
	UserStatusSummary    *ManagedDeviceMobileAppConfigurationUserSummary          `json:"userStatusSummary,omitempty"`
	UserStatuses         *[]ManagedDeviceMobileAppConfigurationUserStatus         `json:"userStatuses,omitempty"`
	Version              *int64                                                   `json:"version,omitempty"`
}
