package workflowtriggers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkflowTriggerProperties struct {
	ChangedTime       *string                           `json:"changedTime,omitempty"`
	CreatedTime       *string                           `json:"createdTime,omitempty"`
	LastExecutionTime *string                           `json:"lastExecutionTime,omitempty"`
	NextExecutionTime *string                           `json:"nextExecutionTime,omitempty"`
	ProvisioningState *WorkflowTriggerProvisioningState `json:"provisioningState,omitempty"`
	Recurrence        *WorkflowTriggerRecurrence        `json:"recurrence,omitempty"`
	State             *WorkflowState                    `json:"state,omitempty"`
	Status            *WorkflowStatus                   `json:"status,omitempty"`
	Workflow          *ResourceReference                `json:"workflow,omitempty"`
}
