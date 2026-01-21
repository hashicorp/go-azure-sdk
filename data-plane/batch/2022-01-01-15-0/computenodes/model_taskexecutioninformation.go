package computenodes

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TaskExecutionInformation struct {
	ContainerInfo   *TaskContainerExecutionInformation `json:"containerInfo,omitempty"`
	EndTime         *string                            `json:"endTime,omitempty"`
	ExitCode        *int64                             `json:"exitCode,omitempty"`
	FailureInfo     *TaskFailureInformation            `json:"failureInfo,omitempty"`
	LastRequeueTime *string                            `json:"lastRequeueTime,omitempty"`
	LastRetryTime   *string                            `json:"lastRetryTime,omitempty"`
	RequeueCount    int64                              `json:"requeueCount"`
	Result          *TaskExecutionResult               `json:"result,omitempty"`
	RetryCount      int64                              `json:"retryCount"`
	StartTime       *string                            `json:"startTime,omitempty"`
}

func (o *TaskExecutionInformation) GetEndTimeAsTime() (*time.Time, error) {
	if o.EndTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *TaskExecutionInformation) SetEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndTime = &formatted
}

func (o *TaskExecutionInformation) GetLastRequeueTimeAsTime() (*time.Time, error) {
	if o.LastRequeueTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastRequeueTime, "2006-01-02T15:04:05Z07:00")
}

func (o *TaskExecutionInformation) SetLastRequeueTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastRequeueTime = &formatted
}

func (o *TaskExecutionInformation) GetLastRetryTimeAsTime() (*time.Time, error) {
	if o.LastRetryTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastRetryTime, "2006-01-02T15:04:05Z07:00")
}

func (o *TaskExecutionInformation) SetLastRetryTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastRetryTime = &formatted
}

func (o *TaskExecutionInformation) GetStartTimeAsTime() (*time.Time, error) {
	if o.StartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *TaskExecutionInformation) SetStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTime = &formatted
}
