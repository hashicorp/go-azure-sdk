package jobs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobSchedule struct {
	Enabled   *bool            `json:"enabled,omitempty"`
	EndTime   *string          `json:"endTime,omitempty"`
	Interval  *string          `json:"interval,omitempty"`
	StartTime *string          `json:"startTime,omitempty"`
	Type      *JobScheduleType `json:"type,omitempty"`
}
