package backups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerBackupProperties struct {
	BackupType    *Origin `json:"backupType,omitempty"`
	CompletedTime *string `json:"completedTime,omitempty"`
	Source        *string `json:"source,omitempty"`
}
