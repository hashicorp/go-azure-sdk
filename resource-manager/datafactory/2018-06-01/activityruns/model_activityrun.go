package activityruns

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
