package devicecapacityinfos

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterGpuCapacity struct {
	GpuFreeUnitsCount                *int64  `json:"gpuFreeUnitsCount,omitempty"`
	GpuReservedForFailoverUnitsCount *int64  `json:"gpuReservedForFailoverUnitsCount,omitempty"`
	GpuTotalUnitsCount               *int64  `json:"gpuTotalUnitsCount,omitempty"`
	GpuType                          *string `json:"gpuType,omitempty"`
	GpuUsedUnitsCount                *int64  `json:"gpuUsedUnitsCount,omitempty"`
}
