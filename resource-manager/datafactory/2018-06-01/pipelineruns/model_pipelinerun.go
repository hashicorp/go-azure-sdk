package pipelineruns

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
