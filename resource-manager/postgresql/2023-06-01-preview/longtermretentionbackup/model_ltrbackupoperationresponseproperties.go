package longtermretentionbackup

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LtrBackupOperationResponseProperties struct {
	BackupMetadata         *string         `json:"backupMetadata,omitempty"`
	BackupName             *string         `json:"backupName,omitempty"`
	DataTransferredInBytes *int64          `json:"dataTransferredInBytes,omitempty"`
	DatasourceSizeInBytes  *int64          `json:"datasourceSizeInBytes,omitempty"`
	EndTime                *string         `json:"endTime,omitempty"`
	ErrorCode              *string         `json:"errorCode,omitempty"`
	ErrorMessage           *string         `json:"errorMessage,omitempty"`
	PercentComplete        *float64        `json:"percentComplete,omitempty"`
	StartTime              string          `json:"startTime"`
	Status                 ExecutionStatus `json:"status"`
}
