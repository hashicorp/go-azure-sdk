package customoperation

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigratePostgreSqlAzureDbForPostgreSqlSyncTaskInput struct {
	EncryptedKeyForSecureFields *string                                                  `json:"encryptedKeyForSecureFields,omitempty"`
	SelectedDatabases           []MigratePostgreSqlAzureDbForPostgreSqlSyncDatabaseInput `json:"selectedDatabases"`
	SourceConnectionInfo        PostgreSqlConnectionInfo                                 `json:"sourceConnectionInfo"`
	StartedOn                   *string                                                  `json:"startedOn,omitempty"`
	TargetConnectionInfo        PostgreSqlConnectionInfo                                 `json:"targetConnectionInfo"`
}

func (o *MigratePostgreSqlAzureDbForPostgreSqlSyncTaskInput) GetStartedOnAsTime() (*time.Time, error) {
	if o.StartedOn == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartedOn, "2006-01-02T15:04:05Z07:00")
}

func (o *MigratePostgreSqlAzureDbForPostgreSqlSyncTaskInput) SetStartedOnAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartedOn = &formatted
}
