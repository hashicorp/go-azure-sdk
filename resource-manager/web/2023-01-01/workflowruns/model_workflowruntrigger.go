package workflowruns

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkflowRunTrigger struct {
	Code              *string         `json:"code,omitempty"`
	Correlation       *Correlation    `json:"correlation,omitempty"`
	EndTime           *string         `json:"endTime,omitempty"`
	Error             *interface{}    `json:"error,omitempty"`
	Inputs            *interface{}    `json:"inputs,omitempty"`
	InputsLink        *ContentLink    `json:"inputsLink,omitempty"`
	Name              *string         `json:"name,omitempty"`
	Outputs           *interface{}    `json:"outputs,omitempty"`
	OutputsLink       *ContentLink    `json:"outputsLink,omitempty"`
	ScheduledTime     *string         `json:"scheduledTime,omitempty"`
	StartTime         *string         `json:"startTime,omitempty"`
	Status            *WorkflowStatus `json:"status,omitempty"`
	TrackedProperties *interface{}    `json:"trackedProperties,omitempty"`
	TrackingId        *string         `json:"trackingId,omitempty"`
}

func (o *WorkflowRunTrigger) GetEndTimeAsTime() (*time.Time, error) {
	if o.EndTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *WorkflowRunTrigger) GetScheduledTimeAsTime() (*time.Time, error) {
	if o.ScheduledTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ScheduledTime, "2006-01-02T15:04:05Z07:00")
}

func (o *WorkflowRunTrigger) GetStartTimeAsTime() (*time.Time, error) {
	if o.StartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartTime, "2006-01-02T15:04:05Z07:00")
}
