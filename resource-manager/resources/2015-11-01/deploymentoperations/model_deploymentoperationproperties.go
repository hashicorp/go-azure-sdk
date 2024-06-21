package deploymentoperations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeploymentOperationProperties struct {
	ProvisioningState *string         `json:"provisioningState,omitempty"`
	StatusCode        *string         `json:"statusCode,omitempty"`
	StatusMessage     *interface{}    `json:"statusMessage,omitempty"`
	TargetResource    *TargetResource `json:"targetResource,omitempty"`
	Timestamp         *string         `json:"timestamp,omitempty"`
}
