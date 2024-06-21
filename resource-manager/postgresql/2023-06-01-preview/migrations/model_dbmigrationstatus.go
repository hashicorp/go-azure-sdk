package migrations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DbMigrationStatus struct {
	AppliedChanges          *int64            `json:"appliedChanges,omitempty"`
	CdcDeleteCounter        *int64            `json:"cdcDeleteCounter,omitempty"`
	CdcInsertCounter        *int64            `json:"cdcInsertCounter,omitempty"`
	CdcUpdateCounter        *int64            `json:"cdcUpdateCounter,omitempty"`
	DatabaseName            *string           `json:"databaseName,omitempty"`
	EndedOn                 *string           `json:"endedOn,omitempty"`
	FullLoadCompletedTables *int64            `json:"fullLoadCompletedTables,omitempty"`
	FullLoadErroredTables   *int64            `json:"fullLoadErroredTables,omitempty"`
	FullLoadLoadingTables   *int64            `json:"fullLoadLoadingTables,omitempty"`
	FullLoadQueuedTables    *int64            `json:"fullLoadQueuedTables,omitempty"`
	IncomingChanges         *int64            `json:"incomingChanges,omitempty"`
	Latency                 *int64            `json:"latency,omitempty"`
	Message                 *string           `json:"message,omitempty"`
	MigrationOperation      *string           `json:"migrationOperation,omitempty"`
	MigrationState          *MigrationDbState `json:"migrationState,omitempty"`
	StartedOn               *string           `json:"startedOn,omitempty"`
}
