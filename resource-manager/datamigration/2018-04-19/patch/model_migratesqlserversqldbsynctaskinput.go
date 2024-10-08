package patch

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrateSqlServerSqlDbSyncTaskInput struct {
	SelectedDatabases    []MigrateSqlServerSqlDbSyncDatabaseInput `json:"selectedDatabases"`
	SourceConnectionInfo SqlConnectionInfo                        `json:"sourceConnectionInfo"`
	TargetConnectionInfo SqlConnectionInfo                        `json:"targetConnectionInfo"`
	ValidationOptions    *MigrationValidationOptions              `json:"validationOptions,omitempty"`
}
