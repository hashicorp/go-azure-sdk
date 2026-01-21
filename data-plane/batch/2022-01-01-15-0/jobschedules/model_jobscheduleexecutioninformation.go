package jobschedules

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobScheduleExecutionInformation struct {
	EndTime     *string    `json:"endTime,omitempty"`
	NextRunTime *string    `json:"nextRunTime,omitempty"`
	RecentJob   *RecentJob `json:"recentJob,omitempty"`
}

func (o *JobScheduleExecutionInformation) GetEndTimeAsTime() (*time.Time, error) {
	if o.EndTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *JobScheduleExecutionInformation) SetEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndTime = &formatted
}

func (o *JobScheduleExecutionInformation) GetNextRunTimeAsTime() (*time.Time, error) {
	if o.NextRunTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.NextRunTime, "2006-01-02T15:04:05Z07:00")
}

func (o *JobScheduleExecutionInformation) SetNextRunTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.NextRunTime = &formatted
}
