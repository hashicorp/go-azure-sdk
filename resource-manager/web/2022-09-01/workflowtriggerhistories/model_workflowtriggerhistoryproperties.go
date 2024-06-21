package workflowtriggerhistories

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkflowTriggerHistoryProperties struct {
	Code          *string            `json:"code,omitempty"`
	Correlation   *Correlation       `json:"correlation,omitempty"`
	EndTime       *string            `json:"endTime,omitempty"`
	Error         *interface{}       `json:"error,omitempty"`
	Fired         *bool              `json:"fired,omitempty"`
	InputsLink    *ContentLink       `json:"inputsLink,omitempty"`
	OutputsLink   *ContentLink       `json:"outputsLink,omitempty"`
	Run           *ResourceReference `json:"run,omitempty"`
	ScheduledTime *string            `json:"scheduledTime,omitempty"`
	StartTime     *string            `json:"startTime,omitempty"`
	Status        *WorkflowStatus    `json:"status,omitempty"`
	TrackingId    *string            `json:"trackingId,omitempty"`
}
