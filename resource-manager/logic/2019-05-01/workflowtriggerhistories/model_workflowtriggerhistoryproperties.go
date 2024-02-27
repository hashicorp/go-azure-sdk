package workflowtriggerhistories

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkflowTriggerHistoryProperties struct {
	Code          *string            `json:"code,omitempty"`
	Correlation   *Correlation       `json:"correlation,omitempty"`
	EndTime       *string            `json:"endTime,omitempty"`
	Error         *interface{}       `json:"error,omitempty"`
	Fired         *bool              `json:"fired,omitempty"`
	InputsLink    *ContentLink       `json:"inputsLink,omitempty"`
	OutputsLink   *ContentLink       `json:"outputsLink,omitempty"`
	Run           *ResourceReference `json:"run,omitempty"`
	ScheduledTime *string            `json:"scheduledTime,omitempty"`
	StartTime     *string            `json:"startTime,omitempty"`
	Status        *WorkflowStatus    `json:"status,omitempty"`
	TrackingId    *string            `json:"trackingId,omitempty"`
}

func (o *WorkflowTriggerHistoryProperties) GetEndTimeAsTime() (*time.Time, error) {
	if o.EndTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *WorkflowTriggerHistoryProperties) SetEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndTime = &formatted
}

func (o *WorkflowTriggerHistoryProperties) GetScheduledTimeAsTime() (*time.Time, error) {
	if o.ScheduledTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ScheduledTime, "2006-01-02T15:04:05Z07:00")
}

func (o *WorkflowTriggerHistoryProperties) SetScheduledTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ScheduledTime = &formatted
}

func (o *WorkflowTriggerHistoryProperties) GetStartTimeAsTime() (*time.Time, error) {
	if o.StartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *WorkflowTriggerHistoryProperties) SetStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTime = &formatted
}
