package workspacemanagerassignments

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobItem struct {
	Errors        *[]Error `json:"errors,omitempty"`
	ExecutionTime *string  `json:"executionTime,omitempty"`
	ResourceId    *string  `json:"resourceId,omitempty"`
	Status        *Status  `json:"status,omitempty"`
}

func (o *JobItem) GetExecutionTimeAsTime() (*time.Time, error) {
	if o.ExecutionTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ExecutionTime, "2006-01-02T15:04:05Z07:00")
}

func (o *JobItem) SetExecutionTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ExecutionTime = &formatted
}
