package tasks

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProjectTaskProperties interface {
}

// RawProjectTaskPropertiesImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawProjectTaskPropertiesImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalProjectTaskPropertiesImplementation(input []byte) (ProjectTaskProperties, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ProjectTaskProperties into map[string]interface: %+v", err)
	}

	value, ok := temp["taskType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "ConnectToSource.PostgreSql.Sync") {
		var out ConnectToSourcePostgreSqlSyncTaskProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConnectToSourcePostgreSqlSyncTaskProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ConnectToSource.SqlServer.Sync") {
		var out ConnectToSourceSqlServerSyncTaskProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConnectToSourceSqlServerSyncTaskProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ConnectToSource.SqlServer") {
		var out ConnectToSourceSqlServerTaskProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConnectToSourceSqlServerTaskProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ConnectToTarget.AzureDbForMySql") {
		var out ConnectToTargetAzureDbForMySqlTaskProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConnectToTargetAzureDbForMySqlTaskProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ConnectToTarget.AzureDbForPostgreSql.Sync") {
		var out ConnectToTargetAzureDbForPostgreSqlSyncTaskProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConnectToTargetAzureDbForPostgreSqlSyncTaskProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ConnectToTarget.SqlDb") {
		var out ConnectToTargetSqlDbTaskProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConnectToTargetSqlDbTaskProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ConnectToTarget.AzureSqlDbMI.Sync.LRS") {
		var out ConnectToTargetSqlMISyncTaskProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConnectToTargetSqlMISyncTaskProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ConnectToTarget.AzureSqlDbMI") {
		var out ConnectToTargetSqlMITaskProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConnectToTargetSqlMITaskProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ConnectToTarget.SqlDb.Sync") {
		var out ConnectToTargetSqlSqlDbSyncTaskProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConnectToTargetSqlSqlDbSyncTaskProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "GetTDECertificates.Sql") {
		var out GetTdeCertificatesSqlTaskProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GetTdeCertificatesSqlTaskProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "GetUserTables.AzureSqlDb.Sync") {
		var out GetUserTablesSqlSyncTaskProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GetUserTablesSqlSyncTaskProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "GetUserTables.Sql") {
		var out GetUserTablesSqlTaskProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GetUserTablesSqlTaskProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Migrate.MySql.AzureDbForMySql.Sync") {
		var out MigrateMySqlAzureDbForMySqlSyncTaskProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MigrateMySqlAzureDbForMySqlSyncTaskProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Migrate.PostgreSql.AzureDbForPostgreSql.Sync") {
		var out MigratePostgreSqlAzureDbForPostgreSqlSyncTaskProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MigratePostgreSqlAzureDbForPostgreSqlSyncTaskProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Migrate.SqlServer.AzureSqlDb.Sync") {
		var out MigrateSqlServerSqlDbSyncTaskProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MigrateSqlServerSqlDbSyncTaskProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Migrate.SqlServer.SqlDb") {
		var out MigrateSqlServerSqlDbTaskProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MigrateSqlServerSqlDbTaskProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Migrate.SqlServer.AzureSqlDbMI.Sync.LRS") {
		var out MigrateSqlServerSqlMISyncTaskProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MigrateSqlServerSqlMISyncTaskProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Migrate.SqlServer.AzureSqlDbMI") {
		var out MigrateSqlServerSqlMITaskProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MigrateSqlServerSqlMITaskProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ValidateMigrationInput.SqlServer.SqlDb.Sync") {
		var out ValidateMigrationInputSqlServerSqlDbSyncTaskProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ValidateMigrationInputSqlServerSqlDbSyncTaskProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ValidateMigrationInput.SqlServer.AzureSqlDbMI.Sync.LRS") {
		var out ValidateMigrationInputSqlServerSqlMISyncTaskProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ValidateMigrationInputSqlServerSqlMISyncTaskProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ValidateMigrationInput.SqlServer.AzureSqlDbMI") {
		var out ValidateMigrationInputSqlServerSqlMITaskProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ValidateMigrationInputSqlServerSqlMITaskProperties: %+v", err)
		}
		return out, nil
	}

	out := RawProjectTaskPropertiesImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
