package sqlpools

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlPoolResourceProperties struct {
	Collation                  *string             `json:"collation,omitempty"`
	CreateMode                 *CreateMode         `json:"createMode,omitempty"`
	CreationDate               *string             `json:"creationDate,omitempty"`
	MaxSizeBytes               *int64              `json:"maxSizeBytes,omitempty"`
	ProvisioningState          *string             `json:"provisioningState,omitempty"`
	RecoverableDatabaseId      *string             `json:"recoverableDatabaseId,omitempty"`
	RestorePointInTime         *string             `json:"restorePointInTime,omitempty"`
	SourceDatabaseDeletionDate *string             `json:"sourceDatabaseDeletionDate,omitempty"`
	SourceDatabaseId           *string             `json:"sourceDatabaseId,omitempty"`
	Status                     *string             `json:"status,omitempty"`
	StorageAccountType         *StorageAccountType `json:"storageAccountType,omitempty"`
}
