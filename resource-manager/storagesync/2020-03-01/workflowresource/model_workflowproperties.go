package workflowresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkflowProperties struct {
	LastOperationId *string             `json:"lastOperationId,omitempty"`
	LastStepName    *string             `json:"lastStepName,omitempty"`
	Operation       *OperationDirection `json:"operation,omitempty"`
	Status          *WorkflowStatus     `json:"status,omitempty"`
	Steps           *string             `json:"steps,omitempty"`
}
