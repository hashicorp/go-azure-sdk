package manageddatabaserestoredetails

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

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

func (o *ManagedDatabaseRestoreDetailsBackupSetProperties) GetRestoreFinishedTimestampUtcAsTime() (*time.Time, error) {
	if o.RestoreFinishedTimestampUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.RestoreFinishedTimestampUtc, "2006-01-02T15:04:05Z07:00")
}

func (o *ManagedDatabaseRestoreDetailsBackupSetProperties) GetRestoreStartedTimestampUtcAsTime() (*time.Time, error) {
	if o.RestoreStartedTimestampUtc == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.RestoreStartedTimestampUtc, "2006-01-02T15:04:05Z07:00")
}
