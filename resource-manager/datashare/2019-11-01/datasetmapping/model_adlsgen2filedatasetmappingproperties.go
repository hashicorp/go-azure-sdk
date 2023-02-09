package datasetmapping

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ADLSGen2FileDataSetMappingProperties struct {
	DataSetId            string                `json:"dataSetId"`
	DataSetMappingStatus *DataSetMappingStatus `json:"dataSetMappingStatus,omitempty"`
	FilePath             string                `json:"filePath"`
	FileSystem           string                `json:"fileSystem"`
	OutputType           *OutputType           `json:"outputType,omitempty"`
	ProvisioningState    *ProvisioningState    `json:"provisioningState,omitempty"`
	ResourceGroup        string                `json:"resourceGroup"`
	StorageAccountName   string                `json:"storageAccountName"`
	SubscriptionId       string                `json:"subscriptionId"`
}
