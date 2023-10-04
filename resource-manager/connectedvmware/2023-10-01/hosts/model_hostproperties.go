package hosts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type HostProperties struct {
	CpuMhz               *int64             `json:"cpuMhz,omitempty"`
	CustomResourceName   *string            `json:"customResourceName,omitempty"`
	DatastoreIds         *[]string          `json:"datastoreIds,omitempty"`
	InventoryItemId      *string            `json:"inventoryItemId,omitempty"`
	MemorySizeGB         *int64             `json:"memorySizeGB,omitempty"`
	MoName               *string            `json:"moName,omitempty"`
	MoRefId              *string            `json:"moRefId,omitempty"`
	NetworkIds           *[]string          `json:"networkIds,omitempty"`
	OverallCPUUsageMHz   *int64             `json:"overallCpuUsageMHz,omitempty"`
	OverallMemoryUsageGB *int64             `json:"overallMemoryUsageGB,omitempty"`
	ProvisioningState    *ProvisioningState `json:"provisioningState,omitempty"`
	Statuses             *[]ResourceStatus  `json:"statuses,omitempty"`
	Uuid                 *string            `json:"uuid,omitempty"`
	VCenterId            *string            `json:"vCenterId,omitempty"`
}
