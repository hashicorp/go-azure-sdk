package backuppolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HourlySchedule struct {
	Interval                *int64  `json:"interval,omitempty"`
	ScheduleWindowDuration  *int64  `json:"scheduleWindowDuration,omitempty"`
	ScheduleWindowStartTime *string `json:"scheduleWindowStartTime,omitempty"`
}
