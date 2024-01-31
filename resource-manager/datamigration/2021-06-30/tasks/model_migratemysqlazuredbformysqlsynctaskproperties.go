package tasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProjectTaskProperties = MigrateMySqlAzureDbForMySqlSyncTaskProperties{}

type MigrateMySqlAzureDbForMySqlSyncTaskProperties struct {
	Input  *MigrateMySqlAzureDbForMySqlSyncTaskInput    `json:"input,omitempty"`
	Output *[]MigrateMySqlAzureDbForMySqlSyncTaskOutput `json:"output,omitempty"`

	// Fields inherited from ProjectTaskProperties
	ClientData *map[string]string   `json:"clientData,omitempty"`
	Commands   *[]CommandProperties `json:"commands,omitempty"`
	Errors     *[]ODataError        `json:"errors,omitempty"`
	State      *TaskState           `json:"state,omitempty"`
}

var _ json.Marshaler = MigrateMySqlAzureDbForMySqlSyncTaskProperties{}

func (s MigrateMySqlAzureDbForMySqlSyncTaskProperties) MarshalJSON() ([]byte, error) {
	type wrapper MigrateMySqlAzureDbForMySqlSyncTaskProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MigrateMySqlAzureDbForMySqlSyncTaskProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MigrateMySqlAzureDbForMySqlSyncTaskProperties: %+v", err)
	}
	decoded["taskType"] = "Migrate.MySql.AzureDbForMySql.Sync"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MigrateMySqlAzureDbForMySqlSyncTaskProperties: %+v", err)
	}

	return encoded, nil
}
