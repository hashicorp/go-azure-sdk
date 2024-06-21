package schedules

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduleProperties struct {
	CreatedDate          *string               `json:"createdDate,omitempty"`
	DailyRecurrence      *DayDetails           `json:"dailyRecurrence,omitempty"`
	HourlyRecurrence     *HourDetails          `json:"hourlyRecurrence,omitempty"`
	NotificationSettings *NotificationSettings `json:"notificationSettings,omitempty"`
	ProvisioningState    *string               `json:"provisioningState,omitempty"`
	Status               *EnableStatus         `json:"status,omitempty"`
	TargetResourceId     *string               `json:"targetResourceId,omitempty"`
	TaskType             *string               `json:"taskType,omitempty"`
	TimeZoneId           *string               `json:"timeZoneId,omitempty"`
	UniqueIdentifier     *string               `json:"uniqueIdentifier,omitempty"`
	WeeklyRecurrence     *WeekDetails          `json:"weeklyRecurrence,omitempty"`
}
