package pipelineruns

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PipelineRun struct {
	DurationInMs  *int64                `json:"durationInMs,omitempty"`
	InvokedBy     *PipelineRunInvokedBy `json:"invokedBy,omitempty"`
	IsLatest      *bool                 `json:"isLatest,omitempty"`
	LastUpdated   *string               `json:"lastUpdated,omitempty"`
	Message       *string               `json:"message,omitempty"`
	Parameters    *map[string]string    `json:"parameters,omitempty"`
	PipelineName  *string               `json:"pipelineName,omitempty"`
	RunDimensions *map[string]string    `json:"runDimensions,omitempty"`
	RunEnd        *string               `json:"runEnd,omitempty"`
	RunGroupId    *string               `json:"runGroupId,omitempty"`
	RunId         *string               `json:"runId,omitempty"`
	RunStart      *string               `json:"runStart,omitempty"`
	Status        *string               `json:"status,omitempty"`
}

func (o *PipelineRun) GetLastUpdatedAsTime() (*time.Time, error) {
	if o.LastUpdated == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastUpdated, "2006-01-02T15:04:05Z07:00")
}

func (o *PipelineRun) GetRunEndAsTime() (*time.Time, error) {
	if o.RunEnd == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.RunEnd, "2006-01-02T15:04:05Z07:00")
}

func (o *PipelineRun) GetRunStartAsTime() (*time.Time, error) {
	if o.RunStart == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.RunStart, "2006-01-02T15:04:05Z07:00")
}
