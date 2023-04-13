package deploymentoperations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeploymentOperation struct {
	Id          *string                        `json:"id,omitempty"`
	OperationId *string                        `json:"operationId,omitempty"`
	Properties  *DeploymentOperationProperties `json:"properties,omitempty"`
}
