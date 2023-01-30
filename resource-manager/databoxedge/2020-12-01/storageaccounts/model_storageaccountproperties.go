package storageaccounts

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type StorageAccountProperties struct {
	BlobEndpoint               *string               `json:"blobEndpoint,omitempty"`
	ContainerCount             *int64                `json:"containerCount,omitempty"`
	DataPolicy                 DataPolicy            `json:"dataPolicy"`
	Description                *string               `json:"description,omitempty"`
	StorageAccountCredentialId *string               `json:"storageAccountCredentialId,omitempty"`
	StorageAccountStatus       *StorageAccountStatus `json:"storageAccountStatus,omitempty"`
}
