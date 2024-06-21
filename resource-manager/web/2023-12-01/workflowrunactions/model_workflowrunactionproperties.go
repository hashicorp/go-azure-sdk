package workflowrunactions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkflowRunActionProperties struct {
	Code              *string               `json:"code,omitempty"`
	Correlation       *RunActionCorrelation `json:"correlation,omitempty"`
	EndTime           *string               `json:"endTime,omitempty"`
	Error             *interface{}          `json:"error,omitempty"`
	InputsLink        *ContentLink          `json:"inputsLink,omitempty"`
	OutputsLink       *ContentLink          `json:"outputsLink,omitempty"`
	RetryHistory      *[]RetryHistory       `json:"retryHistory,omitempty"`
	StartTime         *string               `json:"startTime,omitempty"`
	Status            *WorkflowStatus       `json:"status,omitempty"`
	TrackedProperties *interface{}          `json:"trackedProperties,omitempty"`
	TrackingId        *string               `json:"trackingId,omitempty"`
}
