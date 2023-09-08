package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AospDeviceOwnerDeviceConfiguration struct {
	AppsBlockInstallFromUnknownSources             *bool                                                   `json:"appsBlockInstallFromUnknownSources,omitempty"`
	Assignments                                    *[]DeviceConfigurationAssignment                        `json:"assignments,omitempty"`
	BluetoothBlockConfiguration                    *bool                                                   `json:"bluetoothBlockConfiguration,omitempty"`
	BluetoothBlocked                               *bool                                                   `json:"bluetoothBlocked,omitempty"`
	CameraBlocked                                  *bool                                                   `json:"cameraBlocked,omitempty"`
	CreatedDateTime                                *string                                                 `json:"createdDateTime,omitempty"`
	Description                                    *string                                                 `json:"description,omitempty"`
	DeviceManagementApplicabilityRuleDeviceMode    *DeviceManagementApplicabilityRuleDeviceMode            `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
	DeviceManagementApplicabilityRuleOsEdition     *DeviceManagementApplicabilityRuleOsEdition             `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
	DeviceManagementApplicabilityRuleOsVersion     *DeviceManagementApplicabilityRuleOsVersion             `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
	DeviceSettingStateSummaries                    *[]SettingStateDeviceSummary                            `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                           *DeviceConfigurationDeviceOverview                      `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                                 *[]DeviceConfigurationDeviceStatus                      `json:"deviceStatuses,omitempty"`
	DisplayName                                    *string                                                 `json:"displayName,omitempty"`
	FactoryResetBlocked                            *bool                                                   `json:"factoryResetBlocked,omitempty"`
	GroupAssignments                               *[]DeviceConfigurationGroupAssignment                   `json:"groupAssignments,omitempty"`
	Id                                             *string                                                 `json:"id,omitempty"`
	LastModifiedDateTime                           *string                                                 `json:"lastModifiedDateTime,omitempty"`
	ODataType                                      *string                                                 `json:"@odata.type,omitempty"`
	PasswordMinimumLength                          *int64                                                  `json:"passwordMinimumLength,omitempty"`
	PasswordMinutesOfInactivityBeforeScreenTimeout *int64                                                  `json:"passwordMinutesOfInactivityBeforeScreenTimeout,omitempty"`
	PasswordRequiredType                           *AospDeviceOwnerDeviceConfigurationPasswordRequiredType `json:"passwordRequiredType,omitempty"`
	PasswordSignInFailureCountBeforeFactoryReset   *int64                                                  `json:"passwordSignInFailureCountBeforeFactoryReset,omitempty"`
	RoleScopeTagIds                                *[]string                                               `json:"roleScopeTagIds,omitempty"`
	ScreenCaptureBlocked                           *bool                                                   `json:"screenCaptureBlocked,omitempty"`
	SecurityAllowDebuggingFeatures                 *bool                                                   `json:"securityAllowDebuggingFeatures,omitempty"`
	StorageBlockExternalMedia                      *bool                                                   `json:"storageBlockExternalMedia,omitempty"`
	StorageBlockUsbFileTransfer                    *bool                                                   `json:"storageBlockUsbFileTransfer,omitempty"`
	SupportsScopeTags                              *bool                                                   `json:"supportsScopeTags,omitempty"`
	UserStatusOverview                             *DeviceConfigurationUserOverview                        `json:"userStatusOverview,omitempty"`
	UserStatuses                                   *[]DeviceConfigurationUserStatus                        `json:"userStatuses,omitempty"`
	Version                                        *int64                                                  `json:"version,omitempty"`
	WifiBlockEditConfigurations                    *bool                                                   `json:"wifiBlockEditConfigurations,omitempty"`
}
