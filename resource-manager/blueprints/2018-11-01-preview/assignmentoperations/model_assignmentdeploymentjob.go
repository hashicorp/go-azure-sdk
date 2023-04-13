package assignmentoperations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AssignmentDeploymentJob struct {
	Action     *string                          `json:"action,omitempty"`
	History    *[]AssignmentDeploymentJobResult `json:"history,omitempty"`
	JobId      *string                          `json:"jobId,omitempty"`
	JobState   *string                          `json:"jobState,omitempty"`
	Kind       *string                          `json:"kind,omitempty"`
	RequestUri *string                          `json:"requestUri,omitempty"`
	Result     *AssignmentDeploymentJobResult   `json:"result,omitempty"`
}
