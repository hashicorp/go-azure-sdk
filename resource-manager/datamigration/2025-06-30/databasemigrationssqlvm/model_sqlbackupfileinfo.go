package databasemigrationssqlvm

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlBackupFileInfo struct {
	CopyDuration         *int64   `json:"copyDuration,omitempty"`
	CopyThroughput       *float64 `json:"copyThroughput,omitempty"`
	DataRead             *int64   `json:"dataRead,omitempty"`
	DataWritten          *int64   `json:"dataWritten,omitempty"`
	FamilySequenceNumber *int64   `json:"familySequenceNumber,omitempty"`
	FileName             *string  `json:"fileName,omitempty"`
	Status               *string  `json:"status,omitempty"`
	TotalSize            *int64   `json:"totalSize,omitempty"`
}
