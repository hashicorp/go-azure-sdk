package jobs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobExecutionProperties struct {
	EndTime   *string                   `json:"endTime,omitempty"`
	StartTime *string                   `json:"startTime,omitempty"`
	Status    *JobExecutionRunningState `json:"status,omitempty"`
	Template  *JobExecutionTemplate     `json:"template,omitempty"`
}
