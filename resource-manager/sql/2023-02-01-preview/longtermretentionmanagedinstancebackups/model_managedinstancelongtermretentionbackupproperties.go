package longtermretentionmanagedinstancebackups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedInstanceLongTermRetentionBackupProperties struct {
	BackupExpirationTime      *string                  `json:"backupExpirationTime,omitempty"`
	BackupStorageRedundancy   *BackupStorageRedundancy `json:"backupStorageRedundancy,omitempty"`
	BackupTime                *string                  `json:"backupTime,omitempty"`
	DatabaseDeletionTime      *string                  `json:"databaseDeletionTime,omitempty"`
	DatabaseName              *string                  `json:"databaseName,omitempty"`
	ManagedInstanceCreateTime *string                  `json:"managedInstanceCreateTime,omitempty"`
	ManagedInstanceName       *string                  `json:"managedInstanceName,omitempty"`
}
