package workflowruns

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkflowRunTrigger struct {
	Code              *string         `json:"code,omitempty"`
	Correlation       *Correlation    `json:"correlation,omitempty"`
	EndTime           *string         `json:"endTime,omitempty"`
	Error             *interface{}    `json:"error,omitempty"`
	Inputs            *interface{}    `json:"inputs,omitempty"`
	InputsLink        *ContentLink    `json:"inputsLink,omitempty"`
	Name              *string         `json:"name,omitempty"`
	Outputs           *interface{}    `json:"outputs,omitempty"`
	OutputsLink       *ContentLink    `json:"outputsLink,omitempty"`
	ScheduledTime     *string         `json:"scheduledTime,omitempty"`
	StartTime         *string         `json:"startTime,omitempty"`
	Status            *WorkflowStatus `json:"status,omitempty"`
	TrackedProperties *interface{}    `json:"trackedProperties,omitempty"`
	TrackingId        *string         `json:"trackingId,omitempty"`
}
