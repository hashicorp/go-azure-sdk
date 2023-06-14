package databases

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatabaseProperties struct {
	CatalogCollation                  *CatalogCollationType `json:"catalogCollation,omitempty"`
	Collation                         *string               `json:"collation,omitempty"`
	CreateMode                        *CreateMode           `json:"createMode,omitempty"`
	CreationDate                      *string               `json:"creationDate,omitempty"`
	CurrentServiceObjectiveName       *string               `json:"currentServiceObjectiveName,omitempty"`
	DatabaseId                        *string               `json:"databaseId,omitempty"`
	DefaultSecondaryLocation          *string               `json:"defaultSecondaryLocation,omitempty"`
	ElasticPoolId                     *string               `json:"elasticPoolId,omitempty"`
	FailoverGroupId                   *string               `json:"failoverGroupId,omitempty"`
	LongTermRetentionBackupResourceId *string               `json:"longTermRetentionBackupResourceId,omitempty"`
	MaxSizeBytes                      *int64                `json:"maxSizeBytes,omitempty"`
	RecoverableDatabaseId             *string               `json:"recoverableDatabaseId,omitempty"`
	RecoveryServicesRecoveryPointId   *string               `json:"recoveryServicesRecoveryPointId,omitempty"`
	RestorableDroppedDatabaseId       *string               `json:"restorableDroppedDatabaseId,omitempty"`
	RestorePointInTime                *string               `json:"restorePointInTime,omitempty"`
	SampleName                        *SampleName           `json:"sampleName,omitempty"`
	SourceDatabaseDeletionDate        *string               `json:"sourceDatabaseDeletionDate,omitempty"`
	SourceDatabaseId                  *string               `json:"sourceDatabaseId,omitempty"`
	Status                            *DatabaseStatus       `json:"status,omitempty"`
	ZoneRedundant                     *bool                 `json:"zoneRedundant,omitempty"`
}

func (o *DatabaseProperties) GetCreationDateAsTime() (*time.Time, error) {
	if o.CreationDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreationDate, "2006-01-02T15:04:05Z07:00")
}

func (o *DatabaseProperties) SetCreationDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreationDate = &formatted
}

func (o *DatabaseProperties) GetRestorePointInTimeAsTime() (*time.Time, error) {
	if o.RestorePointInTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.RestorePointInTime, "2006-01-02T15:04:05Z07:00")
}

func (o *DatabaseProperties) SetRestorePointInTimeAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.RestorePointInTime = &formatted
}

func (o *DatabaseProperties) GetSourceDatabaseDeletionDateAsTime() (*time.Time, error) {
	if o.SourceDatabaseDeletionDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.SourceDatabaseDeletionDate, "2006-01-02T15:04:05Z07:00")
}

func (o *DatabaseProperties) SetSourceDatabaseDeletionDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.SourceDatabaseDeletionDate = &formatted
}
