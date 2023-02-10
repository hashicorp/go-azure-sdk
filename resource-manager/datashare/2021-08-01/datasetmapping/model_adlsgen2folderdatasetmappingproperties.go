package datasetmapping

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ADLSGen2FolderDataSetMappingProperties struct {
	DataSetId            string                `json:"dataSetId"`
	DataSetMappingStatus *DataSetMappingStatus `json:"dataSetMappingStatus,omitempty"`
	FileSystem           string                `json:"fileSystem"`
	FolderPath           string                `json:"folderPath"`
	ProvisioningState    *ProvisioningState    `json:"provisioningState,omitempty"`
	ResourceGroup        string                `json:"resourceGroup"`
	StorageAccountName   string                `json:"storageAccountName"`
	SubscriptionId       string                `json:"subscriptionId"`
}
