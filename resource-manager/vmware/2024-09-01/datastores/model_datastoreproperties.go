package datastores

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatastoreProperties struct {
	DiskPoolVolume    *DiskPoolVolume             `json:"diskPoolVolume,omitempty"`
	ElasticSanVolume  *ElasticSanVolume           `json:"elasticSanVolume,omitempty"`
	NetAppVolume      *NetAppVolume               `json:"netAppVolume,omitempty"`
	ProvisioningState *DatastoreProvisioningState `json:"provisioningState,omitempty"`
	PureStorageVolume *PureStorageVolume          `json:"pureStorageVolume,omitempty"`
	Status            *DatastoreStatus            `json:"status,omitempty"`
}
