package roles

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type KubernetesRoleCompute struct {
	MemoryInBytes  *int64 `json:"memoryInBytes,omitempty"`
	ProcessorCount *int64 `json:"processorCount,omitempty"`
	VMProfile      string `json:"vmProfile"`
}
