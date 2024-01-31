package oracleazuredbpostgresqlsynctasks

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MigrateOracleAzureDbPostgreSqlSyncTaskOutput interface {
}

// RawMigrateOracleAzureDbPostgreSqlSyncTaskOutputImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawMigrateOracleAzureDbPostgreSqlSyncTaskOutputImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalMigrateOracleAzureDbPostgreSqlSyncTaskOutputImplementation(input []byte) (MigrateOracleAzureDbPostgreSqlSyncTaskOutput, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling MigrateOracleAzureDbPostgreSqlSyncTaskOutput into map[string]interface: %+v", err)
	}

	value, ok := temp["resultType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "DatabaseLevelErrorOutput") {
		var out MigrateOracleAzureDbPostgreSqlSyncTaskOutputDatabaseError
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MigrateOracleAzureDbPostgreSqlSyncTaskOutputDatabaseError: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "DatabaseLevelOutput") {
		var out MigrateOracleAzureDbPostgreSqlSyncTaskOutputDatabaseLevel
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MigrateOracleAzureDbPostgreSqlSyncTaskOutputDatabaseLevel: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ErrorOutput") {
		var out MigrateOracleAzureDbPostgreSqlSyncTaskOutputError
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MigrateOracleAzureDbPostgreSqlSyncTaskOutputError: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "MigrationLevelOutput") {
		var out MigrateOracleAzureDbPostgreSqlSyncTaskOutputMigrationLevel
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MigrateOracleAzureDbPostgreSqlSyncTaskOutputMigrationLevel: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "TableLevelOutput") {
		var out MigrateOracleAzureDbPostgreSqlSyncTaskOutputTableLevel
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MigrateOracleAzureDbPostgreSqlSyncTaskOutputTableLevel: %+v", err)
		}
		return out, nil
	}

	out := RawMigrateOracleAzureDbPostgreSqlSyncTaskOutputImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
