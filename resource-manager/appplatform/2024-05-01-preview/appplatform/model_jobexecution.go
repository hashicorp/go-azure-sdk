package appplatform

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobExecution struct {
	EndTime     *string                   `json:"endTime,omitempty"`
	JobSnapshot *JobResourceProperties    `json:"jobSnapshot,omitempty"`
	Name        *string                   `json:"name,omitempty"`
	StartTime   *string                   `json:"startTime,omitempty"`
	Status      *JobExecutionRunningState `json:"status,omitempty"`
	Template    *JobExecutionTemplate     `json:"template,omitempty"`
}

func (o *JobExecution) GetEndTimeAsTime() (*time.Time, error) {
	if o.EndTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *JobExecution) SetEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndTime = &formatted
}

func (o *JobExecution) GetStartTimeAsTime() (*time.Time, error) {
	if o.StartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *JobExecution) SetStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTime = &formatted
}
