package jobdefinitions

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduleInfo struct {
	CronExpression *string        `json:"cronExpression,omitempty"`
	DaysOfMonth    *[]int64       `json:"daysOfMonth,omitempty"`
	DaysOfWeek     *[]string      `json:"daysOfWeek,omitempty"`
	EndDate        *string        `json:"endDate,omitempty"`
	ExecutionTime  *SchedulerTime `json:"executionTime,omitempty"`
	Frequency      *Frequency     `json:"frequency,omitempty"`
	IsActive       *bool          `json:"isActive,omitempty"`
	RepeatInterval *string        `json:"repeatInterval,omitempty"`
	StartDate      *string        `json:"startDate,omitempty"`
}

func (o *ScheduleInfo) GetEndDateAsTime() (*time.Time, error) {
	if o.EndDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndDate, "2006-01-02T15:04:05Z07:00")
}

func (o *ScheduleInfo) SetEndDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndDate = &formatted
}

func (o *ScheduleInfo) GetStartDateAsTime() (*time.Time, error) {
	if o.StartDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartDate, "2006-01-02T15:04:05Z07:00")
}

func (o *ScheduleInfo) SetStartDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartDate = &formatted
}
