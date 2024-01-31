package migrateschemasqlserversqldbtasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MigrateSchemaSqlServerSqlDbTaskOutput = MigrateSchemaSqlServerSqlDbTaskOutputMigrationLevel{}

type MigrateSchemaSqlServerSqlDbTaskOutputMigrationLevel struct {
	EndedOn                  *string         `json:"endedOn,omitempty"`
	SourceServerBrandVersion *string         `json:"sourceServerBrandVersion,omitempty"`
	SourceServerVersion      *string         `json:"sourceServerVersion,omitempty"`
	StartedOn                *string         `json:"startedOn,omitempty"`
	State                    *MigrationState `json:"state,omitempty"`
	TargetServerBrandVersion *string         `json:"targetServerBrandVersion,omitempty"`
	TargetServerVersion      *string         `json:"targetServerVersion,omitempty"`

	// Fields inherited from MigrateSchemaSqlServerSqlDbTaskOutput
	Id *string `json:"id,omitempty"`
}

var _ json.Marshaler = MigrateSchemaSqlServerSqlDbTaskOutputMigrationLevel{}

func (s MigrateSchemaSqlServerSqlDbTaskOutputMigrationLevel) MarshalJSON() ([]byte, error) {
	type wrapper MigrateSchemaSqlServerSqlDbTaskOutputMigrationLevel
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MigrateSchemaSqlServerSqlDbTaskOutputMigrationLevel: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MigrateSchemaSqlServerSqlDbTaskOutputMigrationLevel: %+v", err)
	}
	decoded["resultType"] = "MigrationLevelOutput"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MigrateSchemaSqlServerSqlDbTaskOutputMigrationLevel: %+v", err)
	}

	return encoded, nil
}
