package migrateschemasqlserversqldbtasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MigrateSchemaSqlServerSqlDbTaskOutput = MigrateSchemaSqlServerSqlDbTaskOutputError{}

type MigrateSchemaSqlServerSqlDbTaskOutputError struct {
	CommandText *string `json:"commandText,omitempty"`
	ErrorText   *string `json:"errorText,omitempty"`

	// Fields inherited from MigrateSchemaSqlServerSqlDbTaskOutput
	Id *string `json:"id,omitempty"`
}

var _ json.Marshaler = MigrateSchemaSqlServerSqlDbTaskOutputError{}

func (s MigrateSchemaSqlServerSqlDbTaskOutputError) MarshalJSON() ([]byte, error) {
	type wrapper MigrateSchemaSqlServerSqlDbTaskOutputError
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MigrateSchemaSqlServerSqlDbTaskOutputError: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MigrateSchemaSqlServerSqlDbTaskOutputError: %+v", err)
	}
	decoded["resultType"] = "SchemaErrorOutput"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MigrateSchemaSqlServerSqlDbTaskOutputError: %+v", err)
	}

	return encoded, nil
}
