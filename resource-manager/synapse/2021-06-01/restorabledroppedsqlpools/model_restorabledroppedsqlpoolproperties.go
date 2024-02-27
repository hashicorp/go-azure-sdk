package restorabledroppedsqlpools

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RestorableDroppedSqlPoolProperties struct {
	CreationDate          *string `json:"creationDate,omitempty"`
	DatabaseName          *string `json:"databaseName,omitempty"`
	DeletionDate          *string `json:"deletionDate,omitempty"`
	EarliestRestoreDate   *string `json:"earliestRestoreDate,omitempty"`
	Edition               *string `json:"edition,omitempty"`
	ElasticPoolName       *string `json:"elasticPoolName,omitempty"`
	MaxSizeBytes          *string `json:"maxSizeBytes,omitempty"`
	ServiceLevelObjective *string `json:"serviceLevelObjective,omitempty"`
}

func (o *RestorableDroppedSqlPoolProperties) GetCreationDateAsTime() (*time.Time, error) {
	if o.CreationDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.CreationDate, "2006-01-02T15:04:05Z07:00")
}

func (o *RestorableDroppedSqlPoolProperties) SetCreationDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.CreationDate = &formatted
}

func (o *RestorableDroppedSqlPoolProperties) GetDeletionDateAsTime() (*time.Time, error) {
	if o.DeletionDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.DeletionDate, "2006-01-02T15:04:05Z07:00")
}

func (o *RestorableDroppedSqlPoolProperties) SetDeletionDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.DeletionDate = &formatted
}

func (o *RestorableDroppedSqlPoolProperties) GetEarliestRestoreDateAsTime() (*time.Time, error) {
	if o.EarliestRestoreDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EarliestRestoreDate, "2006-01-02T15:04:05Z07:00")
}

func (o *RestorableDroppedSqlPoolProperties) SetEarliestRestoreDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EarliestRestoreDate = &formatted
}
