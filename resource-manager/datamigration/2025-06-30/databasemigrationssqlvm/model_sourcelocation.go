package databasemigrationssqlvm

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SourceLocation struct {
	AzureBlob       *AzureBlob    `json:"azureBlob,omitempty"`
	FileShare       *SqlFileShare `json:"fileShare,omitempty"`
	FileStorageType *string       `json:"fileStorageType,omitempty"`
}
