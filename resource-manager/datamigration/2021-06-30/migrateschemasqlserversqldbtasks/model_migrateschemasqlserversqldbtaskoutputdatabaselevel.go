package migrateschemasqlserversqldbtasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MigrateSchemaSqlServerSqlDbTaskOutput = MigrateSchemaSqlServerSqlDbTaskOutputDatabaseLevel{}

type MigrateSchemaSqlServerSqlDbTaskOutputDatabaseLevel struct {
	DatabaseErrorResultPrefix    *string               `json:"databaseErrorResultPrefix,omitempty"`
	DatabaseName                 *string               `json:"databaseName,omitempty"`
	EndedOn                      *string               `json:"endedOn,omitempty"`
	FileId                       *string               `json:"fileId,omitempty"`
	NumberOfFailedOperations     *int64                `json:"numberOfFailedOperations,omitempty"`
	NumberOfSuccessfulOperations *int64                `json:"numberOfSuccessfulOperations,omitempty"`
	SchemaErrorResultPrefix      *string               `json:"schemaErrorResultPrefix,omitempty"`
	Stage                        *SchemaMigrationStage `json:"stage,omitempty"`
	StartedOn                    *string               `json:"startedOn,omitempty"`
	State                        *MigrationState       `json:"state,omitempty"`

	// Fields inherited from MigrateSchemaSqlServerSqlDbTaskOutput
	Id *string `json:"id,omitempty"`
}

var _ json.Marshaler = MigrateSchemaSqlServerSqlDbTaskOutputDatabaseLevel{}

func (s MigrateSchemaSqlServerSqlDbTaskOutputDatabaseLevel) MarshalJSON() ([]byte, error) {
	type wrapper MigrateSchemaSqlServerSqlDbTaskOutputDatabaseLevel
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MigrateSchemaSqlServerSqlDbTaskOutputDatabaseLevel: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MigrateSchemaSqlServerSqlDbTaskOutputDatabaseLevel: %+v", err)
	}
	decoded["resultType"] = "DatabaseLevelOutput"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MigrateSchemaSqlServerSqlDbTaskOutputDatabaseLevel: %+v", err)
	}

	return encoded, nil
}
