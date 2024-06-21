package longtermretentionbackups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LongTermRetentionBackupProperties struct {
	BackupExpirationTime             *string                  `json:"backupExpirationTime,omitempty"`
	BackupStorageAccessTier          *BackupStorageAccessTier `json:"backupStorageAccessTier,omitempty"`
	BackupStorageRedundancy          *BackupStorageRedundancy `json:"backupStorageRedundancy,omitempty"`
	BackupTime                       *string                  `json:"backupTime,omitempty"`
	DatabaseDeletionTime             *string                  `json:"databaseDeletionTime,omitempty"`
	DatabaseName                     *string                  `json:"databaseName,omitempty"`
	IsBackupImmutable                *bool                    `json:"isBackupImmutable,omitempty"`
	RequestedBackupStorageRedundancy *BackupStorageRedundancy `json:"requestedBackupStorageRedundancy,omitempty"`
	ServerCreateTime                 *string                  `json:"serverCreateTime,omitempty"`
	ServerName                       *string                  `json:"serverName,omitempty"`
}
