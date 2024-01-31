package tasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProjectTaskProperties = MigrateMySqlAzureDbForMySqlOfflineTaskProperties{}

type MigrateMySqlAzureDbForMySqlOfflineTaskProperties struct {
	Input  *MigrateMySqlAzureDbForMySqlOfflineTaskInput    `json:"input,omitempty"`
	Output *[]MigrateMySqlAzureDbForMySqlOfflineTaskOutput `json:"output,omitempty"`

	// Fields inherited from ProjectTaskProperties
	ClientData *map[string]string   `json:"clientData,omitempty"`
	Commands   *[]CommandProperties `json:"commands,omitempty"`
	Errors     *[]ODataError        `json:"errors,omitempty"`
	State      *TaskState           `json:"state,omitempty"`
}

var _ json.Marshaler = MigrateMySqlAzureDbForMySqlOfflineTaskProperties{}

func (s MigrateMySqlAzureDbForMySqlOfflineTaskProperties) MarshalJSON() ([]byte, error) {
	type wrapper MigrateMySqlAzureDbForMySqlOfflineTaskProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MigrateMySqlAzureDbForMySqlOfflineTaskProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MigrateMySqlAzureDbForMySqlOfflineTaskProperties: %+v", err)
	}
	decoded["taskType"] = "Migrate.MySql.AzureDbForMySql"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MigrateMySqlAzureDbForMySqlOfflineTaskProperties: %+v", err)
	}

	return encoded, nil
}
