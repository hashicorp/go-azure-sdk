package backupjobs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DpmJobTaskDetails struct {
	Duration  *string `json:"duration,omitempty"`
	EndTime   *string `json:"endTime,omitempty"`
	StartTime *string `json:"startTime,omitempty"`
	Status    *string `json:"status,omitempty"`
	TaskId    *string `json:"taskId,omitempty"`
}
