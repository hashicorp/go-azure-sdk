package containerappsrevisionreplicas

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicaProperties struct {
	Containers          *[]ReplicaContainer              `json:"containers,omitempty"`
	CreatedTime         *string                          `json:"createdTime,omitempty"`
	InitContainers      *[]ReplicaContainer              `json:"initContainers,omitempty"`
	RunningState        *ContainerAppReplicaRunningState `json:"runningState,omitempty"`
	RunningStateDetails *string                          `json:"runningStateDetails,omitempty"`
}
