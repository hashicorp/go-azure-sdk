package workflowrunoperations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkflowRunProperties struct {
	Code          *string                             `json:"code,omitempty"`
	Correlation   *Correlation                        `json:"correlation,omitempty"`
	CorrelationId *string                             `json:"correlationId,omitempty"`
	EndTime       *string                             `json:"endTime,omitempty"`
	Error         *interface{}                        `json:"error,omitempty"`
	Outputs       *map[string]WorkflowOutputParameter `json:"outputs,omitempty"`
	Response      *WorkflowRunTrigger                 `json:"response,omitempty"`
	StartTime     *string                             `json:"startTime,omitempty"`
	Status        *WorkflowStatus                     `json:"status,omitempty"`
	Trigger       *WorkflowRunTrigger                 `json:"trigger,omitempty"`
	WaitEndTime   *string                             `json:"waitEndTime,omitempty"`
	Workflow      *ResourceReference                  `json:"workflow,omitempty"`
}
