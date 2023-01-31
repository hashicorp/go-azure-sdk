package roles

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KubernetesClusterInfo struct {
	EtcdInfo *EtcdInfo   `json:"etcdInfo,omitempty"`
	Nodes    *[]NodeInfo `json:"nodes,omitempty"`
	Version  string      `json:"version"`
}
