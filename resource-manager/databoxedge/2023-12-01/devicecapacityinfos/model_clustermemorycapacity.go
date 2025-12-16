package devicecapacityinfos

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterMemoryCapacity struct {
	ClusterFailoverMemoryMb      *float64 `json:"clusterFailoverMemoryMb,omitempty"`
	ClusterFragmentationMemoryMb *float64 `json:"clusterFragmentationMemoryMb,omitempty"`
	ClusterFreeMemoryMb          *float64 `json:"clusterFreeMemoryMb,omitempty"`
	ClusterHypervReserveMemoryMb *float64 `json:"clusterHypervReserveMemoryMb,omitempty"`
	ClusterInfraVMMemoryMb       *float64 `json:"clusterInfraVmMemoryMb,omitempty"`
	ClusterMemoryUsedByVMsMb     *float64 `json:"clusterMemoryUsedByVmsMb,omitempty"`
	ClusterNonFailoverVMMb       *float64 `json:"clusterNonFailoverVmMb,omitempty"`
	ClusterTotalMemoryMb         *float64 `json:"clusterTotalMemoryMb,omitempty"`
	ClusterUsedMemoryMb          *float64 `json:"clusterUsedMemoryMb,omitempty"`
}
