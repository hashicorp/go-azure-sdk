package longtermretentionmanagedinstancebackups

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ManagedInstanceLongTermRetentionBackupProperties struct {
	BackupExpirationTime      *string `json:"backupExpirationTime,omitempty"`
	BackupTime                *string `json:"backupTime,omitempty"`
	DatabaseDeletionTime      *string `json:"databaseDeletionTime,omitempty"`
	DatabaseName              *string `json:"databaseName,omitempty"`
	ManagedInstanceCreateTime *string `json:"managedInstanceCreateTime,omitempty"`
	ManagedInstanceName       *string `json:"managedInstanceName,omitempty"`
}

func (o *ManagedInstanceLongTermRetentionBackupProperties) GetBackupExpirationTimeAsTime() (*time.Time, error) {
	if o.BackupExpirationTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.BackupExpirationTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ManagedInstanceLongTermRetentionBackupProperties) SetBackupExpirationTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.BackupExpirationTime = &formatted
}

func (o *ManagedInstanceLongTermRetentionBackupProperties) GetBackupTimeAsTime() (*time.Time, error) {
	if o.BackupTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.BackupTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ManagedInstanceLongTermRetentionBackupProperties) SetBackupTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.BackupTime = &formatted
}

func (o *ManagedInstanceLongTermRetentionBackupProperties) GetDatabaseDeletionTimeAsTime() (*time.Time, error) {
	if o.DatabaseDeletionTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.DatabaseDeletionTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ManagedInstanceLongTermRetentionBackupProperties) SetDatabaseDeletionTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.DatabaseDeletionTime = &formatted
}

func (o *ManagedInstanceLongTermRetentionBackupProperties) GetManagedInstanceCreateTimeAsTime() (*time.Time, error) {
	if o.ManagedInstanceCreateTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ManagedInstanceCreateTime, "2006-01-02T15:04:05Z07:00")
}

func (o *ManagedInstanceLongTermRetentionBackupProperties) SetManagedInstanceCreateTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ManagedInstanceCreateTime = &formatted
}
