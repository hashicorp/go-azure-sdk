package manageddatabaserestoredetails

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDatabaseRestoreDetailsBackupSetProperties struct {
	BackupSizeMB                *int64  `json:"backupSizeMB,omitempty"`
	FirstStripeName             *string `json:"firstStripeName,omitempty"`
	NumberOfStripes             *int64  `json:"numberOfStripes,omitempty"`
	RestoreFinishedTimestampUtc *string `json:"restoreFinishedTimestampUtc,omitempty"`
	RestoreStartedTimestampUtc  *string `json:"restoreStartedTimestampUtc,omitempty"`
	Status                      *string `json:"status,omitempty"`
}
