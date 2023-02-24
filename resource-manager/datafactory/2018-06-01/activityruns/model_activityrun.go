package activityruns

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ActivityRun struct {
	ActivityName      *string      `json:"activityName,omitempty"`
	ActivityRunEnd    *string      `json:"activityRunEnd,omitempty"`
	ActivityRunId     *string      `json:"activityRunId,omitempty"`
	ActivityRunStart  *string      `json:"activityRunStart,omitempty"`
	ActivityType      *string      `json:"activityType,omitempty"`
	DurationInMs      *int64       `json:"durationInMs,omitempty"`
	Error             *interface{} `json:"error,omitempty"`
	Input             *interface{} `json:"input,omitempty"`
	LinkedServiceName *string      `json:"linkedServiceName,omitempty"`
	Output            *interface{} `json:"output,omitempty"`
	PipelineName      *string      `json:"pipelineName,omitempty"`
	PipelineRunId     *string      `json:"pipelineRunId,omitempty"`
	Status            *string      `json:"status,omitempty"`
}

func (o *ActivityRun) GetActivityRunEndAsTime() (*time.Time, error) {
	if o.ActivityRunEnd == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ActivityRunEnd, "2006-01-02T15:04:05Z07:00")
}

func (o *ActivityRun) SetActivityRunEndAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ActivityRunEnd = &formatted
}

func (o *ActivityRun) GetActivityRunStartAsTime() (*time.Time, error) {
	if o.ActivityRunStart == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ActivityRunStart, "2006-01-02T15:04:05Z07:00")
}

func (o *ActivityRun) SetActivityRunStartAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ActivityRunStart = &formatted
}
