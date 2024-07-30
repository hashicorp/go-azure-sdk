package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Windows10TeamGeneralConfiguration struct {
	Assignments                            *[]DeviceConfigurationAssignment                                  `json:"assignments,omitempty"`
	AzureOperationalInsightsBlockTelemetry *bool                                                             `json:"azureOperationalInsightsBlockTelemetry,omitempty"`
	AzureOperationalInsightsWorkspaceId    *string                                                           `json:"azureOperationalInsightsWorkspaceId,omitempty"`
	AzureOperationalInsightsWorkspaceKey   *string                                                           `json:"azureOperationalInsightsWorkspaceKey,omitempty"`
	ConnectAppBlockAutoLaunch              *bool                                                             `json:"connectAppBlockAutoLaunch,omitempty"`
	CreatedDateTime                        *string                                                           `json:"createdDateTime,omitempty"`
	Description                            *string                                                           `json:"description,omitempty"`
	DeviceSettingStateSummaries            *[]SettingStateDeviceSummary                                      `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                   *DeviceConfigurationDeviceOverview                                `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                         *[]DeviceConfigurationDeviceStatus                                `json:"deviceStatuses,omitempty"`
	DisplayName                            *string                                                           `json:"displayName,omitempty"`
	Id                                     *string                                                           `json:"id,omitempty"`
	LastModifiedDateTime                   *string                                                           `json:"lastModifiedDateTime,omitempty"`
	MaintenanceWindowBlocked               *bool                                                             `json:"maintenanceWindowBlocked,omitempty"`
	MaintenanceWindowDurationInHours       *int64                                                            `json:"maintenanceWindowDurationInHours,omitempty"`
	MaintenanceWindowStartTime             *string                                                           `json:"maintenanceWindowStartTime,omitempty"`
	MiracastBlocked                        *bool                                                             `json:"miracastBlocked,omitempty"`
	MiracastChannel                        *Windows10TeamGeneralConfigurationMiracastChannel                 `json:"miracastChannel,omitempty"`
	MiracastRequirePin                     *bool                                                             `json:"miracastRequirePin,omitempty"`
	ODataType                              *string                                                           `json:"@odata.type,omitempty"`
	SettingsBlockMyMeetingsAndFiles        *bool                                                             `json:"settingsBlockMyMeetingsAndFiles,omitempty"`
	SettingsBlockSessionResume             *bool                                                             `json:"settingsBlockSessionResume,omitempty"`
	SettingsBlockSigninSuggestions         *bool                                                             `json:"settingsBlockSigninSuggestions,omitempty"`
	SettingsDefaultVolume                  *int64                                                            `json:"settingsDefaultVolume,omitempty"`
	SettingsScreenTimeoutInMinutes         *int64                                                            `json:"settingsScreenTimeoutInMinutes,omitempty"`
	SettingsSessionTimeoutInMinutes        *int64                                                            `json:"settingsSessionTimeoutInMinutes,omitempty"`
	SettingsSleepTimeoutInMinutes          *int64                                                            `json:"settingsSleepTimeoutInMinutes,omitempty"`
	UserStatusOverview                     *DeviceConfigurationUserOverview                                  `json:"userStatusOverview,omitempty"`
	UserStatuses                           *[]DeviceConfigurationUserStatus                                  `json:"userStatuses,omitempty"`
	Version                                *int64                                                            `json:"version,omitempty"`
	WelcomeScreenBackgroundImageUrl        *string                                                           `json:"welcomeScreenBackgroundImageUrl,omitempty"`
	WelcomeScreenBlockAutomaticWakeUp      *bool                                                             `json:"welcomeScreenBlockAutomaticWakeUp,omitempty"`
	WelcomeScreenMeetingInformation        *Windows10TeamGeneralConfigurationWelcomeScreenMeetingInformation `json:"welcomeScreenMeetingInformation,omitempty"`
}
