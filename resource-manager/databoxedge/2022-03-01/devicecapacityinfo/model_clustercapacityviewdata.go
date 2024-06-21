package devicecapacityinfo

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterCapacityViewData struct {
	Fqdn                        *string                `json:"fqdn,omitempty"`
	GpuCapacity                 *ClusterGpuCapacity    `json:"gpuCapacity,omitempty"`
	LastRefreshedTime           *string                `json:"lastRefreshedTime,omitempty"`
	MemoryCapacity              *ClusterMemoryCapacity `json:"memoryCapacity,omitempty"`
	TotalProvisionedNonHpnCores *int64                 `json:"totalProvisionedNonHpnCores,omitempty"`
}
