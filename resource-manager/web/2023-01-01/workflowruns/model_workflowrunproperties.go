package workflowruns

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkflowRunProperties struct {
	Code          *string                             `json:"code,omitempty"`
	Correlation   *Correlation                        `json:"correlation,omitempty"`
	CorrelationId *string                             `json:"correlationId,omitempty"`
	EndTime       *string                             `json:"endTime,omitempty"`
	Error         *interface{}                        `json:"error,omitempty"`
	Outputs       *map[string]WorkflowOutputParameter `json:"outputs,omitempty"`
	Response      *WorkflowRunTrigger                 `json:"response,omitempty"`
	StartTime     *string                             `json:"startTime,omitempty"`
	Status        *WorkflowStatus                     `json:"status,omitempty"`
	Trigger       *WorkflowRunTrigger                 `json:"trigger,omitempty"`
	WaitEndTime   *string                             `json:"waitEndTime,omitempty"`
	Workflow      *ResourceReference                  `json:"workflow,omitempty"`
}

func (o *WorkflowRunProperties) GetEndTimeAsTime() (*time.Time, error) {
	if o.EndTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *WorkflowRunProperties) GetStartTimeAsTime() (*time.Time, error) {
	if o.StartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *WorkflowRunProperties) GetWaitEndTimeAsTime() (*time.Time, error) {
	if o.WaitEndTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.WaitEndTime, "2006-01-02T15:04:05Z07:00")
}
