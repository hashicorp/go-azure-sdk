package schedule

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduleCreateOrUpdateProperties struct {
	AdvancedSchedule *AdvancedSchedule `json:"advancedSchedule,omitempty"`
	Description      *string           `json:"description,omitempty"`
	ExpiryTime       *string           `json:"expiryTime,omitempty"`
	Frequency        ScheduleFrequency `json:"frequency"`
	Interval         *interface{}      `json:"interval,omitempty"`
	StartTime        string            `json:"startTime"`
	TimeZone         *string           `json:"timeZone,omitempty"`
}
