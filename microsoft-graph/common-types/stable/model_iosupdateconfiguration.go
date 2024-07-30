package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IosUpdateConfiguration struct {
	ActiveHoursEnd              *string                                     `json:"activeHoursEnd,omitempty"`
	ActiveHoursStart            *string                                     `json:"activeHoursStart,omitempty"`
	Assignments                 *[]DeviceConfigurationAssignment            `json:"assignments,omitempty"`
	CreatedDateTime             *string                                     `json:"createdDateTime,omitempty"`
	Description                 *string                                     `json:"description,omitempty"`
	DeviceSettingStateSummaries *[]SettingStateDeviceSummary                `json:"deviceSettingStateSummaries,omitempty"`
	DeviceStatusOverview        *DeviceConfigurationDeviceOverview          `json:"deviceStatusOverview,omitempty"`
	DeviceStatuses              *[]DeviceConfigurationDeviceStatus          `json:"deviceStatuses,omitempty"`
	DisplayName                 *string                                     `json:"displayName,omitempty"`
	Id                          *string                                     `json:"id,omitempty"`
	LastModifiedDateTime        *string                                     `json:"lastModifiedDateTime,omitempty"`
	ODataType                   *string                                     `json:"@odata.type,omitempty"`
	ScheduledInstallDays        *IosUpdateConfigurationScheduledInstallDays `json:"scheduledInstallDays,omitempty"`
	UserStatusOverview          *DeviceConfigurationUserOverview            `json:"userStatusOverview,omitempty"`
	UserStatuses                *[]DeviceConfigurationUserStatus            `json:"userStatuses,omitempty"`
	UtcTimeOffsetInMinutes      *int64                                      `json:"utcTimeOffsetInMinutes,omitempty"`
	Version                     *int64                                      `json:"version,omitempty"`
}
