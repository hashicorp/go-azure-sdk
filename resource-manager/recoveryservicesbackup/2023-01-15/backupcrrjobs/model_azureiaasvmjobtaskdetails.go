package backupcrrjobs

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureIaaSVMJobTaskDetails struct {
	Duration             *string  `json:"duration,omitempty"`
	EndTime              *string  `json:"endTime,omitempty"`
	InstanceId           *string  `json:"instanceId,omitempty"`
	ProgressPercentage   *float64 `json:"progressPercentage,omitempty"`
	StartTime            *string  `json:"startTime,omitempty"`
	Status               *string  `json:"status,omitempty"`
	TaskExecutionDetails *string  `json:"taskExecutionDetails,omitempty"`
	TaskId               *string  `json:"taskId,omitempty"`
}

func (o *AzureIaaSVMJobTaskDetails) GetEndTimeAsTime() (*time.Time, error) {
	if o.EndTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndTime, "2006-01-02T15:04:05Z07:00")
}

func (o *AzureIaaSVMJobTaskDetails) SetEndTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndTime = &formatted
}

func (o *AzureIaaSVMJobTaskDetails) GetStartTimeAsTime() (*time.Time, error) {
	if o.StartTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartTime, "2006-01-02T15:04:05Z07:00")
}

func (o *AzureIaaSVMJobTaskDetails) SetStartTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartTime = &formatted
}
