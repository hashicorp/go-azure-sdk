package computenodes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TaskInformation struct {
	ExecutionInfo *TaskExecutionInformation `json:"executionInfo,omitempty"`
	JobId         *string                   `json:"jobId,omitempty"`
	SubtaskId     *int64                    `json:"subtaskId,omitempty"`
	TaskId        *string                   `json:"taskId,omitempty"`
	TaskState     TaskState                 `json:"taskState"`
	TaskURL       *string                   `json:"taskUrl,omitempty"`
}
