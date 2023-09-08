package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsKioskConfiguration struct {
	Assignments                                 *[]DeviceConfigurationAssignment             `json:"assignments,omitempty"`
	CreatedDateTime                             *string                                      `json:"createdDateTime,omitempty"`
	Description                                 *string                                      `json:"description,omitempty"`
	DeviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
	DeviceManagementApplicabilityRuleOsEdition  *DeviceManagementApplicabilityRuleOsEdition  `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
	DeviceManagementApplicabilityRuleOsVersion  *DeviceManagementApplicabilityRuleOsVersion  `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
	DeviceSettingStateSummaries                 *[]SettingStateDeviceSummary                 `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                        *DeviceConfigurationDeviceOverview           `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                              *[]DeviceConfigurationDeviceStatus           `json:"deviceStatuses,omitempty"`
	DisplayName                                 *string                                      `json:"displayName,omitempty"`
	EdgeKioskEnablePublicBrowsing               *bool                                        `json:"edgeKioskEnablePublicBrowsing,omitempty"`
	GroupAssignments                            *[]DeviceConfigurationGroupAssignment        `json:"groupAssignments,omitempty"`
	Id                                          *string                                      `json:"id,omitempty"`
	KioskBrowserBlockedURLs                     *[]string                                    `json:"kioskBrowserBlockedURLs,omitempty"`
	KioskBrowserBlockedUrlExceptions            *[]string                                    `json:"kioskBrowserBlockedUrlExceptions,omitempty"`
	KioskBrowserDefaultUrl                      *string                                      `json:"kioskBrowserDefaultUrl,omitempty"`
	KioskBrowserEnableEndSessionButton          *bool                                        `json:"kioskBrowserEnableEndSessionButton,omitempty"`
	KioskBrowserEnableHomeButton                *bool                                        `json:"kioskBrowserEnableHomeButton,omitempty"`
	KioskBrowserEnableNavigationButtons         *bool                                        `json:"kioskBrowserEnableNavigationButtons,omitempty"`
	KioskBrowserRestartOnIdleTimeInMinutes      *int64                                       `json:"kioskBrowserRestartOnIdleTimeInMinutes,omitempty"`
	KioskProfiles                               *[]WindowsKioskProfile                       `json:"kioskProfiles,omitempty"`
	LastModifiedDateTime                        *string                                      `json:"lastModifiedDateTime,omitempty"`
	ODataType                                   *string                                      `json:"@odata.type,omitempty"`
	RoleScopeTagIds                             *[]string                                    `json:"roleScopeTagIds,omitempty"`
	SupportsScopeTags                           *bool                                        `json:"supportsScopeTags,omitempty"`
	UserStatusOverview                          *DeviceConfigurationUserOverview             `json:"userStatusOverview,omitempty"`
	UserStatuses                                *[]DeviceConfigurationUserStatus             `json:"userStatuses,omitempty"`
	Version                                     *int64                                       `json:"version,omitempty"`
	WindowsKioskForceUpdateSchedule             *WindowsKioskForceUpdateSchedule             `json:"windowsKioskForceUpdateSchedule,omitempty"`
}
