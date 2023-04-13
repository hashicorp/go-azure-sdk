package assignmentoperations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentOperationProperties struct {
	AssignmentState  *string                    `json:"assignmentState,omitempty"`
	BlueprintVersion *string                    `json:"blueprintVersion,omitempty"`
	Deployments      *[]AssignmentDeploymentJob `json:"deployments,omitempty"`
	TimeCreated      *string                    `json:"timeCreated,omitempty"`
	TimeFinished     *string                    `json:"timeFinished,omitempty"`
	TimeStarted      *string                    `json:"timeStarted,omitempty"`
}
