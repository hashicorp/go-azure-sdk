package elasticsan

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ElasticSanProperties struct {
	AvailabilityZones       *Zones              `json:"availabilityZones,omitempty"`
	BaseSizeTiB             int64               `json:"baseSizeTiB"`
	ExtendedCapacitySizeTiB int64               `json:"extendedCapacitySizeTiB"`
	ProvisioningState       *ProvisioningStates `json:"provisioningState,omitempty"`
	Sku                     Sku                 `json:"sku"`
	TotalIops               *int64              `json:"totalIops,omitempty"`
	TotalMBps               *int64              `json:"totalMBps,omitempty"`
	TotalSizeTiB            *int64              `json:"totalSizeTiB,omitempty"`
	TotalVolumeSizeGiB      *int64              `json:"totalVolumeSizeGiB,omitempty"`
	VolumeGroupCount        *int64              `json:"volumeGroupCount,omitempty"`
}
