package workflowrunoperations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type WorkflowOutputParameter struct {
	Description *string        `json:"description,omitempty"`
	Error       *interface{}   `json:"error,omitempty"`
	Metadata    *interface{}   `json:"metadata,omitempty"`
	Type        *ParameterType `json:"type,omitempty"`
	Value       *interface{}   `json:"value,omitempty"`
}
