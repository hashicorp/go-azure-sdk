package tasks

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SubtaskInformation struct {
	ContainerInfo               *TaskContainerExecutionInformation `json:"containerInfo,omitempty"`
	EndTime                     *string                            `json:"endTime,omitempty"`
	ExitCode                    *int64                             `json:"exitCode,omitempty"`
	FailureInfo                 *TaskFailureInformation            `json:"failureInfo,omitempty"`
	Id                          *int64                             `json:"id,omitempty"`
	NodeInfo                    *ComputeNodeInformation            `json:"nodeInfo,omitempty"`
	PreviousState               *SubtaskState                      `json:"previousState,omitempty"`
	PreviousStateTransitionTime *string                            `json:"previousStateTransitionTime,omitempty"`
	Result                      *TaskExecutionResult               `json:"result,omitempty"`
	StartTime                   *string                            `json:"startTime,omitempty"`
	State                       *SubtaskState                      `json:"state,omitempty"`
	StateTransitionTime         *string                            `json:"stateTransitionTime,omitempty"`
}

func (o *SubtaskInformation) GetEndTimeAsTime() (*time.Time, error) {
	if o.EndTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *SubtaskInformation) SetEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndTime = &formatted
}

func (o *SubtaskInformation) GetPreviousStateTransitionTimeAsTime() (*time.Time, error) {
	if o.PreviousStateTransitionTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.PreviousStateTransitionTime, "2006-01-02T15:04:05Z07:00")
}

func (o *SubtaskInformation) SetPreviousStateTransitionTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.PreviousStateTransitionTime = &formatted
}

func (o *SubtaskInformation) GetStartTimeAsTime() (*time.Time, error) {
	if o.StartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *SubtaskInformation) SetStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTime = &formatted
}

func (o *SubtaskInformation) GetStateTransitionTimeAsTime() (*time.Time, error) {
	if o.StateTransitionTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StateTransitionTime, "2006-01-02T15:04:05Z07:00")
}

func (o *SubtaskInformation) SetStateTransitionTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StateTransitionTime = &formatted
}
