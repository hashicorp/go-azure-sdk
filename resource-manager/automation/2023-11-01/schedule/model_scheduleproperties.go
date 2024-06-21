package schedule

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduleProperties struct {
	AdvancedSchedule        *AdvancedSchedule  `json:"advancedSchedule,omitempty"`
	CreationTime            *string            `json:"creationTime,omitempty"`
	Description             *string            `json:"description,omitempty"`
	ExpiryTime              *string            `json:"expiryTime,omitempty"`
	ExpiryTimeOffsetMinutes *float64           `json:"expiryTimeOffsetMinutes,omitempty"`
	Frequency               *ScheduleFrequency `json:"frequency,omitempty"`
	Interval                *interface{}       `json:"interval,omitempty"`
	IsEnabled               *bool              `json:"isEnabled,omitempty"`
	LastModifiedTime        *string            `json:"lastModifiedTime,omitempty"`
	NextRun                 *string            `json:"nextRun,omitempty"`
	NextRunOffsetMinutes    *float64           `json:"nextRunOffsetMinutes,omitempty"`
	StartTime               *string            `json:"startTime,omitempty"`
	StartTimeOffsetMinutes  *float64           `json:"startTimeOffsetMinutes,omitempty"`
	TimeZone                *string            `json:"timeZone,omitempty"`
}
