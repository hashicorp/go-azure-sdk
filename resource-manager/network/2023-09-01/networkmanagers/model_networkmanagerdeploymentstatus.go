package networkmanagers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkManagerDeploymentStatus struct {
	CommitTime       *string            `json:"commitTime,omitempty"`
	ConfigurationIds *[]string          `json:"configurationIds,omitempty"`
	DeploymentStatus *DeploymentStatus  `json:"deploymentStatus,omitempty"`
	DeploymentType   *ConfigurationType `json:"deploymentType,omitempty"`
	ErrorMessage     *string            `json:"errorMessage,omitempty"`
	Region           *string            `json:"region,omitempty"`
}
