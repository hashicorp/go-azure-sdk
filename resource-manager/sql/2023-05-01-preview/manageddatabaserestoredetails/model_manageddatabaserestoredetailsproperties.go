package manageddatabaserestoredetails

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedDatabaseRestoreDetailsProperties struct {
	BlockReason               *string                                                    `json:"blockReason,omitempty"`
	CurrentBackupType         *string                                                    `json:"currentBackupType,omitempty"`
	CurrentRestorePlanSizeMB  *int64                                                     `json:"currentRestorePlanSizeMB,omitempty"`
	CurrentRestoredSizeMB     *int64                                                     `json:"currentRestoredSizeMB,omitempty"`
	CurrentRestoringFileName  *string                                                    `json:"currentRestoringFileName,omitempty"`
	DiffBackupSets            *[]ManagedDatabaseRestoreDetailsBackupSetProperties        `json:"diffBackupSets,omitempty"`
	FullBackupSets            *[]ManagedDatabaseRestoreDetailsBackupSetProperties        `json:"fullBackupSets,omitempty"`
	LastRestoredFileName      *string                                                    `json:"lastRestoredFileName,omitempty"`
	LastRestoredFileTime      *string                                                    `json:"lastRestoredFileTime,omitempty"`
	LastUploadedFileName      *string                                                    `json:"lastUploadedFileName,omitempty"`
	LastUploadedFileTime      *string                                                    `json:"lastUploadedFileTime,omitempty"`
	LogBackupSets             *[]ManagedDatabaseRestoreDetailsBackupSetProperties        `json:"logBackupSets,omitempty"`
	NumberOfFilesDetected     *int64                                                     `json:"numberOfFilesDetected,omitempty"`
	NumberOfFilesQueued       *int64                                                     `json:"numberOfFilesQueued,omitempty"`
	NumberOfFilesRestored     *int64                                                     `json:"numberOfFilesRestored,omitempty"`
	NumberOfFilesRestoring    *int64                                                     `json:"numberOfFilesRestoring,omitempty"`
	NumberOfFilesSkipped      *int64                                                     `json:"numberOfFilesSkipped,omitempty"`
	NumberOfFilesUnrestorable *int64                                                     `json:"numberOfFilesUnrestorable,omitempty"`
	PercentCompleted          *int64                                                     `json:"percentCompleted,omitempty"`
	Status                    *string                                                    `json:"status,omitempty"`
	Type                      *string                                                    `json:"type,omitempty"`
	UnrestorableFiles         *[]ManagedDatabaseRestoreDetailsUnrestorableFileProperties `json:"unrestorableFiles,omitempty"`
}
