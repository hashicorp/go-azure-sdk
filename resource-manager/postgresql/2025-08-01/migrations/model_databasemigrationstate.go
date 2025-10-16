package migrations

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DatabaseMigrationState struct {
	AppliedChanges          *int64                  `json:"appliedChanges,omitempty"`
	CdcDeleteCounter        *int64                  `json:"cdcDeleteCounter,omitempty"`
	CdcInsertCounter        *int64                  `json:"cdcInsertCounter,omitempty"`
	CdcUpdateCounter        *int64                  `json:"cdcUpdateCounter,omitempty"`
	DatabaseName            *string                 `json:"databaseName,omitempty"`
	EndedOn                 *string                 `json:"endedOn,omitempty"`
	FullLoadCompletedTables *int64                  `json:"fullLoadCompletedTables,omitempty"`
	FullLoadErroredTables   *int64                  `json:"fullLoadErroredTables,omitempty"`
	FullLoadLoadingTables   *int64                  `json:"fullLoadLoadingTables,omitempty"`
	FullLoadQueuedTables    *int64                  `json:"fullLoadQueuedTables,omitempty"`
	IncomingChanges         *int64                  `json:"incomingChanges,omitempty"`
	Latency                 *int64                  `json:"latency,omitempty"`
	Message                 *string                 `json:"message,omitempty"`
	MigrationOperation      *string                 `json:"migrationOperation,omitempty"`
	MigrationState          *MigrationDatabaseState `json:"migrationState,omitempty"`
	StartedOn               *string                 `json:"startedOn,omitempty"`
}

func (o *DatabaseMigrationState) GetEndedOnAsTime() (*time.Time, error) {
	if o.EndedOn == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndedOn, "2006-01-02T15:04:05Z07:00")
}

func (o *DatabaseMigrationState) SetEndedOnAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EndedOn = &formatted
}

func (o *DatabaseMigrationState) GetStartedOnAsTime() (*time.Time, error) {
	if o.StartedOn == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartedOn, "2006-01-02T15:04:05Z07:00")
}

func (o *DatabaseMigrationState) SetStartedOnAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartedOn = &formatted
}
