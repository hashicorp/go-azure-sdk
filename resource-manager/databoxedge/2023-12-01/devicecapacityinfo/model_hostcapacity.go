package devicecapacityinfo

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HostCapacity struct {
	AvailableGpuCount                *int64               `json:"availableGpuCount,omitempty"`
	EffectiveAvailableMemoryMbOnHost *int64               `json:"effectiveAvailableMemoryMbOnHost,omitempty"`
	GpuType                          *string              `json:"gpuType,omitempty"`
	HostName                         *string              `json:"hostName,omitempty"`
	NumaNodesData                    *[]NumaNodeData      `json:"numaNodesData,omitempty"`
	VMUsedMemory                     *map[string]VMMemory `json:"vmUsedMemory,omitempty"`
}
