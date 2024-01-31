package tasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProjectTaskProperties = ValidateOracleAzureDbForPostgreSqlSyncTaskProperties{}

type ValidateOracleAzureDbForPostgreSqlSyncTaskProperties struct {
	Input  *MigrateOracleAzureDbPostgreSqlSyncTaskInput     `json:"input,omitempty"`
	Output *[]ValidateOracleAzureDbPostgreSqlSyncTaskOutput `json:"output,omitempty"`

	// Fields inherited from ProjectTaskProperties
	ClientData *map[string]string   `json:"clientData,omitempty"`
	Commands   *[]CommandProperties `json:"commands,omitempty"`
	Errors     *[]ODataError        `json:"errors,omitempty"`
	State      *TaskState           `json:"state,omitempty"`
}

var _ json.Marshaler = ValidateOracleAzureDbForPostgreSqlSyncTaskProperties{}

func (s ValidateOracleAzureDbForPostgreSqlSyncTaskProperties) MarshalJSON() ([]byte, error) {
	type wrapper ValidateOracleAzureDbForPostgreSqlSyncTaskProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ValidateOracleAzureDbForPostgreSqlSyncTaskProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ValidateOracleAzureDbForPostgreSqlSyncTaskProperties: %+v", err)
	}
	decoded["taskType"] = "Validate.Oracle.AzureDbPostgreSql.Sync"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ValidateOracleAzureDbForPostgreSqlSyncTaskProperties: %+v", err)
	}

	return encoded, nil
}
