package devicecapacityinfo

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NumaNodeData struct {
	EffectiveAvailableMemoryInMb *int64   `json:"effectiveAvailableMemoryInMb,omitempty"`
	FreeVCPUIndexesForHpn        *[]int64 `json:"freeVCpuIndexesForHpn,omitempty"`
	LogicalCoreCountPerCore      *int64   `json:"logicalCoreCountPerCore,omitempty"`
	NumaNodeIndex                *int64   `json:"numaNodeIndex,omitempty"`
	TotalMemoryInMb              *int64   `json:"totalMemoryInMb,omitempty"`
	VCPUIndexesForHpn            *[]int64 `json:"vCpuIndexesForHpn,omitempty"`
	VCPUIndexesForRoot           *[]int64 `json:"vCpuIndexesForRoot,omitempty"`
}
