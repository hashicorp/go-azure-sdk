package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WindowsUpdateForBusinessConfiguration struct {
	AllowWindows11Upgrade                       *bool                                                                  `json:"allowWindows11Upgrade,omitempty"`
	Assignments                                 *[]DeviceConfigurationAssignment                                       `json:"assignments,omitempty"`
	AutoRestartNotificationDismissal            *WindowsUpdateForBusinessConfigurationAutoRestartNotificationDismissal `json:"autoRestartNotificationDismissal,omitempty"`
	AutomaticUpdateMode                         *WindowsUpdateForBusinessConfigurationAutomaticUpdateMode              `json:"automaticUpdateMode,omitempty"`
	BusinessReadyUpdatesOnly                    *WindowsUpdateForBusinessConfigurationBusinessReadyUpdatesOnly         `json:"businessReadyUpdatesOnly,omitempty"`
	CreatedDateTime                             *string                                                                `json:"createdDateTime,omitempty"`
	DeadlineForFeatureUpdatesInDays             *int64                                                                 `json:"deadlineForFeatureUpdatesInDays,omitempty"`
	DeadlineForQualityUpdatesInDays             *int64                                                                 `json:"deadlineForQualityUpdatesInDays,omitempty"`
	DeadlineGracePeriodInDays                   *int64                                                                 `json:"deadlineGracePeriodInDays,omitempty"`
	DeliveryOptimizationMode                    *WindowsUpdateForBusinessConfigurationDeliveryOptimizationMode         `json:"deliveryOptimizationMode,omitempty"`
	Description                                 *string                                                                `json:"description,omitempty"`
	DeviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode                           `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
	DeviceManagementApplicabilityRuleOsEdition  *DeviceManagementApplicabilityRuleOsEdition                            `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
	DeviceManagementApplicabilityRuleOsVersion  *DeviceManagementApplicabilityRuleOsVersion                            `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
	DeviceSettingStateSummaries                 *[]SettingStateDeviceSummary                                           `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                        *DeviceConfigurationDeviceOverview                                     `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                              *[]DeviceConfigurationDeviceStatus                                     `json:"deviceStatuses,omitempty"`
	DisplayName                                 *string                                                                `json:"displayName,omitempty"`
	DriversExcluded                             *bool                                                                  `json:"driversExcluded,omitempty"`
	EngagedRestartDeadlineInDays                *int64                                                                 `json:"engagedRestartDeadlineInDays,omitempty"`
	EngagedRestartSnoozeScheduleInDays          *int64                                                                 `json:"engagedRestartSnoozeScheduleInDays,omitempty"`
	EngagedRestartTransitionScheduleInDays      *int64                                                                 `json:"engagedRestartTransitionScheduleInDays,omitempty"`
	FeatureUpdatesDeferralPeriodInDays          *int64                                                                 `json:"featureUpdatesDeferralPeriodInDays,omitempty"`
	FeatureUpdatesPauseExpiryDateTime           *string                                                                `json:"featureUpdatesPauseExpiryDateTime,omitempty"`
	FeatureUpdatesPauseStartDate                *string                                                                `json:"featureUpdatesPauseStartDate,omitempty"`
	FeatureUpdatesPaused                        *bool                                                                  `json:"featureUpdatesPaused,omitempty"`
	FeatureUpdatesRollbackStartDateTime         *string                                                                `json:"featureUpdatesRollbackStartDateTime,omitempty"`
	FeatureUpdatesRollbackWindowInDays          *int64                                                                 `json:"featureUpdatesRollbackWindowInDays,omitempty"`
	FeatureUpdatesWillBeRolledBack              *bool                                                                  `json:"featureUpdatesWillBeRolledBack,omitempty"`
	GroupAssignments                            *[]DeviceConfigurationGroupAssignment                                  `json:"groupAssignments,omitempty"`
	Id                                          *string                                                                `json:"id,omitempty"`
	InstallationSchedule                        *WindowsUpdateInstallScheduleType                                      `json:"installationSchedule,omitempty"`
	LastModifiedDateTime                        *string                                                                `json:"lastModifiedDateTime,omitempty"`
	MicrosoftUpdateServiceAllowed               *bool                                                                  `json:"microsoftUpdateServiceAllowed,omitempty"`
	ODataType                                   *string                                                                `json:"@odata.type,omitempty"`
	PostponeRebootUntilAfterDeadline            *bool                                                                  `json:"postponeRebootUntilAfterDeadline,omitempty"`
	PrereleaseFeatures                          *WindowsUpdateForBusinessConfigurationPrereleaseFeatures               `json:"prereleaseFeatures,omitempty"`
	QualityUpdatesDeferralPeriodInDays          *int64                                                                 `json:"qualityUpdatesDeferralPeriodInDays,omitempty"`
	QualityUpdatesPauseExpiryDateTime           *string                                                                `json:"qualityUpdatesPauseExpiryDateTime,omitempty"`
	QualityUpdatesPauseStartDate                *string                                                                `json:"qualityUpdatesPauseStartDate,omitempty"`
	QualityUpdatesPaused                        *bool                                                                  `json:"qualityUpdatesPaused,omitempty"`
	QualityUpdatesRollbackStartDateTime         *string                                                                `json:"qualityUpdatesRollbackStartDateTime,omitempty"`
	QualityUpdatesWillBeRolledBack              *bool                                                                  `json:"qualityUpdatesWillBeRolledBack,omitempty"`
	RoleScopeTagIds                             *[]string                                                              `json:"roleScopeTagIds,omitempty"`
	ScheduleImminentRestartWarningInMinutes     *int64                                                                 `json:"scheduleImminentRestartWarningInMinutes,omitempty"`
	ScheduleRestartWarningInHours               *int64                                                                 `json:"scheduleRestartWarningInHours,omitempty"`
	SkipChecksBeforeRestart                     *bool                                                                  `json:"skipChecksBeforeRestart,omitempty"`
	SupportsScopeTags                           *bool                                                                  `json:"supportsScopeTags,omitempty"`
	UpdateNotificationLevel                     *WindowsUpdateForBusinessConfigurationUpdateNotificationLevel          `json:"updateNotificationLevel,omitempty"`
	UpdateWeeks                                 *WindowsUpdateForBusinessConfigurationUpdateWeeks                      `json:"updateWeeks,omitempty"`
	UserPauseAccess                             *WindowsUpdateForBusinessConfigurationUserPauseAccess                  `json:"userPauseAccess,omitempty"`
	UserStatusOverview                          *DeviceConfigurationUserOverview                                       `json:"userStatusOverview,omitempty"`
	UserStatuses                                *[]DeviceConfigurationUserStatus                                       `json:"userStatuses,omitempty"`
	UserWindowsUpdateScanAccess                 *WindowsUpdateForBusinessConfigurationUserWindowsUpdateScanAccess      `json:"userWindowsUpdateScanAccess,omitempty"`
	Version                                     *int64                                                                 `json:"version,omitempty"`
}
