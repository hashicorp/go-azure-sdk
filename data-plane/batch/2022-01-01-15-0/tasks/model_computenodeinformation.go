package tasks

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComputeNodeInformation struct {
	AffinityId           *string `json:"affinityId,omitempty"`
	NodeId               *string `json:"nodeId,omitempty"`
	NodeURL              *string `json:"nodeUrl,omitempty"`
	PoolId               *string `json:"poolId,omitempty"`
	TaskRootDirectory    *string `json:"taskRootDirectory,omitempty"`
	TaskRootDirectoryURL *string `json:"taskRootDirectoryUrl,omitempty"`
}
