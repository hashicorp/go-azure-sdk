package resourcepools

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ResourcePoolProperties struct {
	CpuCapacityMHz     *int64             `json:"cpuCapacityMHz,omitempty"`
	CpuLimitMHz        *int64             `json:"cpuLimitMHz,omitempty"`
	CpuOverallUsageMHz *int64             `json:"cpuOverallUsageMHz,omitempty"`
	CpuReservationMHz  *int64             `json:"cpuReservationMHz,omitempty"`
	CpuSharesLevel     *string            `json:"cpuSharesLevel,omitempty"`
	CustomResourceName *string            `json:"customResourceName,omitempty"`
	DatastoreIds       *[]string          `json:"datastoreIds,omitempty"`
	InventoryItemId    *string            `json:"inventoryItemId,omitempty"`
	MemCapacityGB      *int64             `json:"memCapacityGB,omitempty"`
	MemLimitMB         *int64             `json:"memLimitMB,omitempty"`
	MemOverallUsageGB  *int64             `json:"memOverallUsageGB,omitempty"`
	MemReservationMB   *int64             `json:"memReservationMB,omitempty"`
	MemSharesLevel     *string            `json:"memSharesLevel,omitempty"`
	MoName             *string            `json:"moName,omitempty"`
	MoRefId            *string            `json:"moRefId,omitempty"`
	NetworkIds         *[]string          `json:"networkIds,omitempty"`
	ProvisioningState  *ProvisioningState `json:"provisioningState,omitempty"`
	Statuses           *[]ResourceStatus  `json:"statuses,omitempty"`
	Uuid               *string            `json:"uuid,omitempty"`
	VCenterId          *string            `json:"vCenterId,omitempty"`
}
