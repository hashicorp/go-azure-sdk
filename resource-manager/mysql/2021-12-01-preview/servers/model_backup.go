package servers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Backup struct {
	BackupRetentionDays *int64            `json:"backupRetentionDays,omitempty"`
	EarliestRestoreDate *string           `json:"earliestRestoreDate,omitempty"`
	GeoRedundantBackup  *EnableStatusEnum `json:"geoRedundantBackup,omitempty"`
}
