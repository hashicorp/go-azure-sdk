package triggers

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduleTriggerRecurrence struct {
	EndTime   *string              `json:"endTime,omitempty"`
	Frequency *RecurrenceFrequency `json:"frequency,omitempty"`
	Interval  *int64               `json:"interval,omitempty"`
	Schedule  *RecurrenceSchedule  `json:"schedule,omitempty"`
	StartTime *string              `json:"startTime,omitempty"`
	TimeZone  *string              `json:"timeZone,omitempty"`
}

func (o *ScheduleTriggerRecurrence) GetEndTimeAsTime() (*time.Time, error) {
	if o.EndTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ScheduleTriggerRecurrence) SetEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndTime = &formatted
}

func (o *ScheduleTriggerRecurrence) GetStartTimeAsTime() (*time.Time, error) {
	if o.StartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ScheduleTriggerRecurrence) SetStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTime = &formatted
}
