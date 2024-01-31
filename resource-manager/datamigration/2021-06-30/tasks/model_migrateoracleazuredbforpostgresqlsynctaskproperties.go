package tasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProjectTaskProperties = MigrateOracleAzureDbForPostgreSqlSyncTaskProperties{}

type MigrateOracleAzureDbForPostgreSqlSyncTaskProperties struct {
	Input  *MigrateOracleAzureDbPostgreSqlSyncTaskInput    `json:"input,omitempty"`
	Output *[]MigrateOracleAzureDbPostgreSqlSyncTaskOutput `json:"output,omitempty"`

	// Fields inherited from ProjectTaskProperties
	ClientData *map[string]string   `json:"clientData,omitempty"`
	Commands   *[]CommandProperties `json:"commands,omitempty"`
	Errors     *[]ODataError        `json:"errors,omitempty"`
	State      *TaskState           `json:"state,omitempty"`
}

var _ json.Marshaler = MigrateOracleAzureDbForPostgreSqlSyncTaskProperties{}

func (s MigrateOracleAzureDbForPostgreSqlSyncTaskProperties) MarshalJSON() ([]byte, error) {
	type wrapper MigrateOracleAzureDbForPostgreSqlSyncTaskProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MigrateOracleAzureDbForPostgreSqlSyncTaskProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MigrateOracleAzureDbForPostgreSqlSyncTaskProperties: %+v", err)
	}
	decoded["taskType"] = "Migrate.Oracle.AzureDbForPostgreSql.Sync"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MigrateOracleAzureDbForPostgreSqlSyncTaskProperties: %+v", err)
	}

	return encoded, nil
}
