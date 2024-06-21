package backupjobs

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
