package tasks

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TaskAddResult struct {
	ETag         *string       `json:"eTag,omitempty"`
	Error        *BatchError   `json:"error,omitempty"`
	LastModified *string       `json:"lastModified,omitempty"`
	Location     *string       `json:"location,omitempty"`
	Status       TaskAddStatus `json:"status"`
	TaskId       string        `json:"taskId"`
}

func (o *TaskAddResult) GetLastModifiedAsTime() (*time.Time, error) {
	if o.LastModified == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastModified, "2006-01-02T15:04:05Z07:00")
}

func (o *TaskAddResult) SetLastModifiedAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastModified = &formatted
}
