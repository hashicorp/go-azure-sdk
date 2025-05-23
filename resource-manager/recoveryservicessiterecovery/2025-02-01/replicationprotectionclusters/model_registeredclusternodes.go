package replicationprotectionclusters

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegisteredClusterNodes struct {
	BiosId                  *string `json:"biosId,omitempty"`
	ClusterNodeFqdn         *string `json:"clusterNodeFqdn,omitempty"`
	IsSharedDiskVirtualNode *bool   `json:"isSharedDiskVirtualNode,omitempty"`
	MachineId               *string `json:"machineId,omitempty"`
}
