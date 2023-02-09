package datasetmapping

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SynapseWorkspaceSqlPoolTableDataSetMappingProperties struct {
	DataSetId                              string                `json:"dataSetId"`
	DataSetMappingStatus                   *DataSetMappingStatus `json:"dataSetMappingStatus,omitempty"`
	ProvisioningState                      *ProvisioningState    `json:"provisioningState,omitempty"`
	SynapseWorkspaceSqlPoolTableResourceId string                `json:"synapseWorkspaceSqlPoolTableResourceId"`
}
