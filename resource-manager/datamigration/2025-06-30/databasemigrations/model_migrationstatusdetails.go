package databasemigrations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrationStatusDetails struct {
	ActiveBackupSets            *[]SqlBackupSetInfo `json:"activeBackupSets,omitempty"`
	BlobContainerName           *string             `json:"blobContainerName,omitempty"`
	CompleteRestoreErrorMessage *string             `json:"completeRestoreErrorMessage,omitempty"`
	CurrentRestoringFilename    *string             `json:"currentRestoringFilename,omitempty"`
	FileUploadBlockingErrors    *[]string           `json:"fileUploadBlockingErrors,omitempty"`
	FullBackupSetInfo           *SqlBackupSetInfo   `json:"fullBackupSetInfo,omitempty"`
	InvalidFiles                *[]string           `json:"invalidFiles,omitempty"`
	IsFullBackupRestored        *bool               `json:"isFullBackupRestored,omitempty"`
	LastRestoredBackupSetInfo   *SqlBackupSetInfo   `json:"lastRestoredBackupSetInfo,omitempty"`
	LastRestoredFilename        *string             `json:"lastRestoredFilename,omitempty"`
	MigrationState              *string             `json:"migrationState,omitempty"`
	PendingLogBackupsCount      *int64              `json:"pendingLogBackupsCount,omitempty"`
	RestoreBlockingReason       *string             `json:"restoreBlockingReason,omitempty"`
}
