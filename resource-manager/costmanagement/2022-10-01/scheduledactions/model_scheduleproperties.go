package scheduledactions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduleProperties struct {
	DayOfMonth   *int64            `json:"dayOfMonth,omitempty"`
	DaysOfWeek   *[]DaysOfWeek     `json:"daysOfWeek,omitempty"`
	EndDate      string            `json:"endDate"`
	Frequency    ScheduleFrequency `json:"frequency"`
	HourOfDay    *int64            `json:"hourOfDay,omitempty"`
	StartDate    string            `json:"startDate"`
	WeeksOfMonth *[]WeeksOfMonth   `json:"weeksOfMonth,omitempty"`
}
