package restorabledroppeddatabases

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorableDroppedDatabaseProperties struct {
	BackupStorageRedundancy *BackupStorageRedundancy `json:"backupStorageRedundancy,omitempty"`
	CreationDate            *string                  `json:"creationDate,omitempty"`
	DatabaseName            *string                  `json:"databaseName,omitempty"`
	DeletionDate            *string                  `json:"deletionDate,omitempty"`
	EarliestRestoreDate     *string                  `json:"earliestRestoreDate,omitempty"`
	Keys                    *map[string]DatabaseKey  `json:"keys,omitempty"`
	MaxSizeBytes            *int64                   `json:"maxSizeBytes,omitempty"`
}
