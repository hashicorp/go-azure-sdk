package migrateschemasqlserversqldbtasks

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrateSchemaSqlServerSqlDbTaskOutput interface {
}

// RawMigrateSchemaSqlServerSqlDbTaskOutputImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawMigrateSchemaSqlServerSqlDbTaskOutputImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalMigrateSchemaSqlServerSqlDbTaskOutputImplementation(input []byte) (MigrateSchemaSqlServerSqlDbTaskOutput, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling MigrateSchemaSqlServerSqlDbTaskOutput into map[string]interface: %+v", err)
	}

	value, ok := temp["resultType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "DatabaseLevelOutput") {
		var out MigrateSchemaSqlServerSqlDbTaskOutputDatabaseLevel
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MigrateSchemaSqlServerSqlDbTaskOutputDatabaseLevel: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SchemaErrorOutput") {
		var out MigrateSchemaSqlServerSqlDbTaskOutputError
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MigrateSchemaSqlServerSqlDbTaskOutputError: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "MigrationLevelOutput") {
		var out MigrateSchemaSqlServerSqlDbTaskOutputMigrationLevel
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MigrateSchemaSqlServerSqlDbTaskOutputMigrationLevel: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ErrorOutput") {
		var out MigrateSchemaSqlTaskOutputError
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MigrateSchemaSqlTaskOutputError: %+v", err)
		}
		return out, nil
	}

	out := RawMigrateSchemaSqlServerSqlDbTaskOutputImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
