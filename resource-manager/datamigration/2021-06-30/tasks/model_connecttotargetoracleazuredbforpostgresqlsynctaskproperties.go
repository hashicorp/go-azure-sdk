package tasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProjectTaskProperties = ConnectToTargetOracleAzureDbForPostgreSqlSyncTaskProperties{}

type ConnectToTargetOracleAzureDbForPostgreSqlSyncTaskProperties struct {
	Input  *ConnectToTargetOracleAzureDbForPostgreSqlSyncTaskInput    `json:"input,omitempty"`
	Output *[]ConnectToTargetOracleAzureDbForPostgreSqlSyncTaskOutput `json:"output,omitempty"`

	// Fields inherited from ProjectTaskProperties
	ClientData *map[string]string   `json:"clientData,omitempty"`
	Commands   *[]CommandProperties `json:"commands,omitempty"`
	Errors     *[]ODataError        `json:"errors,omitempty"`
	State      *TaskState           `json:"state,omitempty"`
}

var _ json.Marshaler = ConnectToTargetOracleAzureDbForPostgreSqlSyncTaskProperties{}

func (s ConnectToTargetOracleAzureDbForPostgreSqlSyncTaskProperties) MarshalJSON() ([]byte, error) {
	type wrapper ConnectToTargetOracleAzureDbForPostgreSqlSyncTaskProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ConnectToTargetOracleAzureDbForPostgreSqlSyncTaskProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ConnectToTargetOracleAzureDbForPostgreSqlSyncTaskProperties: %+v", err)
	}
	decoded["taskType"] = "ConnectToTarget.Oracle.AzureDbForPostgreSql.Sync"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ConnectToTargetOracleAzureDbForPostgreSqlSyncTaskProperties: %+v", err)
	}

	return encoded, nil
}
