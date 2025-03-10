package sourcecontrols

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeploymentInfo struct {
	Deployment            *Deployment            `json:"deployment,omitempty"`
	DeploymentFetchStatus *DeploymentFetchStatus `json:"deploymentFetchStatus,omitempty"`
	Message               *string                `json:"message,omitempty"`
}
