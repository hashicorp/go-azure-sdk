package sourcecontrols

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Deployment struct {
	DeploymentId      *string           `json:"deploymentId,omitempty"`
	DeploymentLogsUrl *string           `json:"deploymentLogsUrl,omitempty"`
	DeploymentResult  *DeploymentResult `json:"deploymentResult,omitempty"`
	DeploymentState   *DeploymentState  `json:"deploymentState,omitempty"`
	DeploymentTime    *string           `json:"deploymentTime,omitempty"`
}
