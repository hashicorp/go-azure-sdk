package containerappsrevisionreplicas

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReplicaContainer struct {
	ContainerId       *string `json:"containerId,omitempty"`
	ExecEndpoint      *string `json:"execEndpoint,omitempty"`
	LogStreamEndpoint *string `json:"logStreamEndpoint,omitempty"`
	Name              *string `json:"name,omitempty"`
	Ready             *bool   `json:"ready,omitempty"`
	RestartCount      *int64  `json:"restartCount,omitempty"`
	Started           *bool   `json:"started,omitempty"`
}
