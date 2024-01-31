package migrateschemasqlserversqldbtasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MigrateSchemaSqlServerSqlDbTaskOutput = MigrateSchemaSqlTaskOutputError{}

type MigrateSchemaSqlTaskOutputError struct {
	Error *ReportableException `json:"error,omitempty"`

	// Fields inherited from MigrateSchemaSqlServerSqlDbTaskOutput
	Id *string `json:"id,omitempty"`
}

var _ json.Marshaler = MigrateSchemaSqlTaskOutputError{}

func (s MigrateSchemaSqlTaskOutputError) MarshalJSON() ([]byte, error) {
	type wrapper MigrateSchemaSqlTaskOutputError
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MigrateSchemaSqlTaskOutputError: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MigrateSchemaSqlTaskOutputError: %+v", err)
	}
	decoded["resultType"] = "ErrorOutput"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MigrateSchemaSqlTaskOutputError: %+v", err)
	}

	return encoded, nil
}
