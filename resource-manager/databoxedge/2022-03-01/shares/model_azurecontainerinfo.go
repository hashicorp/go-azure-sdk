package shares

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AzureContainerInfo struct {
	ContainerName              string                   `json:"containerName"`
	DataFormat                 AzureContainerDataFormat `json:"dataFormat"`
	StorageAccountCredentialId string                   `json:"storageAccountCredentialId"`
}
