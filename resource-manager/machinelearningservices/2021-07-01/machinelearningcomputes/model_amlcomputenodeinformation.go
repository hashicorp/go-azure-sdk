package machinelearningcomputes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AmlComputeNodeInformation struct {
	NodeId           *string    `json:"nodeId,omitempty"`
	NodeState        *NodeState `json:"nodeState,omitempty"`
	Port             *float64   `json:"port,omitempty"`
	PrivateIpAddress *string    `json:"privateIpAddress,omitempty"`
	PublicIpAddress  *string    `json:"publicIpAddress,omitempty"`
	RunId            *string    `json:"runId,omitempty"`
}
