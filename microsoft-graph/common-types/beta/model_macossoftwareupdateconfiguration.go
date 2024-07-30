package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MacOSSoftwareUpdateConfiguration struct {
	AllOtherUpdateBehavior                      *MacOSSoftwareUpdateConfigurationAllOtherUpdateBehavior   `json:"allOtherUpdateBehavior,omitempty"`
	Assignments                                 *[]DeviceConfigurationAssignment                          `json:"assignments,omitempty"`
	ConfigDataUpdateBehavior                    *MacOSSoftwareUpdateConfigurationConfigDataUpdateBehavior `json:"configDataUpdateBehavior,omitempty"`
	CreatedDateTime                             *string                                                   `json:"createdDateTime,omitempty"`
	CriticalUpdateBehavior                      *MacOSSoftwareUpdateConfigurationCriticalUpdateBehavior   `json:"criticalUpdateBehavior,omitempty"`
	CustomUpdateTimeWindows                     *[]CustomUpdateTimeWindow                                 `json:"customUpdateTimeWindows,omitempty"`
	Description                                 *string                                                   `json:"description,omitempty"`
	DeviceManagementApplicabilityRuleDeviceMode *DeviceManagementApplicabilityRuleDeviceMode              `json:"deviceManagementApplicabilityRuleDeviceMode,omitempty"`
	DeviceManagementApplicabilityRuleOsEdition  *DeviceManagementApplicabilityRuleOsEdition               `json:"deviceManagementApplicabilityRuleOsEdition,omitempty"`
	DeviceManagementApplicabilityRuleOsVersion  *DeviceManagementApplicabilityRuleOsVersion               `json:"deviceManagementApplicabilityRuleOsVersion,omitempty"`
	DeviceSettingStateSummaries                 *[]SettingStateDeviceSummary                              `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview                        *DeviceConfigurationDeviceOverview                        `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses                              *[]DeviceConfigurationDeviceStatus                        `json:"deviceStatuses,omitempty"`
	DisplayName                                 *string                                                   `json:"displayName,omitempty"`
	FirmwareUpdateBehavior                      *MacOSSoftwareUpdateConfigurationFirmwareUpdateBehavior   `json:"firmwareUpdateBehavior,omitempty"`
	GroupAssignments                            *[]DeviceConfigurationGroupAssignment                     `json:"groupAssignments,omitempty"`
	Id                                          *string                                                   `json:"id,omitempty"`
	LastModifiedDateTime                        *string                                                   `json:"lastModifiedDateTime,omitempty"`
	MaxUserDeferralsCount                       *int64                                                    `json:"maxUserDeferralsCount,omitempty"`
	ODataType                                   *string                                                   `json:"@odata.type,omitempty"`
	Priority                                    *MacOSSoftwareUpdateConfigurationPriority                 `json:"priority,omitempty"`
	RoleScopeTagIds                             *[]string                                                 `json:"roleScopeTagIds,omitempty"`
	SupportsScopeTags                           *bool                                                     `json:"supportsScopeTags,omitempty"`
	UpdateScheduleType                          *MacOSSoftwareUpdateConfigurationUpdateScheduleType       `json:"updateScheduleType,omitempty"`
	UpdateTimeWindowUtcOffsetInMinutes          *int64                                                    `json:"updateTimeWindowUtcOffsetInMinutes,omitempty"`
	UserStatusOverview                          *DeviceConfigurationUserOverview                          `json:"userStatusOverview,omitempty"`
	UserStatuses                                *[]DeviceConfigurationUserStatus                          `json:"userStatuses,omitempty"`
	Version                                     *int64                                                    `json:"version,omitempty"`
}
