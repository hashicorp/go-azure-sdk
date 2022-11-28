package pool

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MountConfiguration struct {
	AzureBlobFileSystemConfiguration *AzureBlobFileSystemConfiguration `json:"azureBlobFileSystemConfiguration"`
	AzureFileShareConfiguration      *AzureFileShareConfiguration      `json:"azureFileShareConfiguration"`
	CifsMountConfiguration           *CIFSMountConfiguration           `json:"cifsMountConfiguration"`
	NfsMountConfiguration            *NFSMountConfiguration            `json:"nfsMountConfiguration"`
}
