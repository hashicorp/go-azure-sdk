package jobs

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
	"github.com/hashicorp/go-azure-helpers/resourcemanager/systemdata"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Job struct {
	EndTime         *string                `json:"endTime,omitempty"`
	Error           *JobErrorDetails       `json:"error,omitempty"`
	Id              *string                `json:"id,omitempty"`
	Name            *string                `json:"name,omitempty"`
	PercentComplete *int64                 `json:"percentComplete,omitempty"`
	Properties      *JobProperties         `json:"properties,omitempty"`
	StartTime       *string                `json:"startTime,omitempty"`
	Status          *JobStatus             `json:"status,omitempty"`
	SystemData      *systemdata.SystemData `json:"systemData,omitempty"`
	Type            *string                `json:"type,omitempty"`
}

func (o *Job) GetEndTimeAsTime() (*time.Time, error) {
	if o.EndTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *Job) SetEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndTime = &formatted
}

func (o *Job) GetStartTimeAsTime() (*time.Time, error) {
	if o.StartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *Job) SetStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTime = &formatted
}
