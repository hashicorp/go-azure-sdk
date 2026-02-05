package accounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PoolNodeCounts struct {
	Dedicated   *NodeCounts `json:"dedicated,omitempty"`
	LowPriority *NodeCounts `json:"lowPriority,omitempty"`
	PoolId      string      `json:"poolId"`
}
