package databasemigrationssqlvm

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SqlBackupSetInfo struct {
	BackupFinishDate   *string              `json:"backupFinishDate,omitempty"`
	BackupSetId        *string              `json:"backupSetId,omitempty"`
	BackupStartDate    *string              `json:"backupStartDate,omitempty"`
	BackupType         *string              `json:"backupType,omitempty"`
	FamilyCount        *int64               `json:"familyCount,omitempty"`
	FirstLSN           *string              `json:"firstLSN,omitempty"`
	HasBackupChecksums *bool                `json:"hasBackupChecksums,omitempty"`
	IgnoreReasons      *[]string            `json:"ignoreReasons,omitempty"`
	IsBackupRestored   *bool                `json:"isBackupRestored,omitempty"`
	LastLSN            *string              `json:"lastLSN,omitempty"`
	ListOfBackupFiles  *[]SqlBackupFileInfo `json:"listOfBackupFiles,omitempty"`
}

func (o *SqlBackupSetInfo) GetBackupFinishDateAsTime() (*time.Time, error) {
	if o.BackupFinishDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.BackupFinishDate, "2006-01-02T15:04:05Z07:00")
}

func (o *SqlBackupSetInfo) SetBackupFinishDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.BackupFinishDate = &formatted
}

func (o *SqlBackupSetInfo) GetBackupStartDateAsTime() (*time.Time, error) {
	if o.BackupStartDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.BackupStartDate, "2006-01-02T15:04:05Z07:00")
}

func (o *SqlBackupSetInfo) SetBackupStartDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.BackupStartDate = &formatted
}
