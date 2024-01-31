package tasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProjectTaskProperties = ValidateMigrationInputSqlServerSqlDbSyncTaskProperties{}

type ValidateMigrationInputSqlServerSqlDbSyncTaskProperties struct {
	Input  *ValidateSyncMigrationInputSqlServerTaskInput    `json:"input,omitempty"`
	Output *[]ValidateSyncMigrationInputSqlServerTaskOutput `json:"output,omitempty"`

	// Fields inherited from ProjectTaskProperties
	Commands *[]CommandProperties `json:"commands,omitempty"`
	Errors   *[]ODataError        `json:"errors,omitempty"`
	State    *TaskState           `json:"state,omitempty"`
}

var _ json.Marshaler = ValidateMigrationInputSqlServerSqlDbSyncTaskProperties{}

func (s ValidateMigrationInputSqlServerSqlDbSyncTaskProperties) MarshalJSON() ([]byte, error) {
	type wrapper ValidateMigrationInputSqlServerSqlDbSyncTaskProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ValidateMigrationInputSqlServerSqlDbSyncTaskProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ValidateMigrationInputSqlServerSqlDbSyncTaskProperties: %+v", err)
	}
	decoded["taskType"] = "ValidateMigrationInput.SqlServer.SqlDb.Sync"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ValidateMigrationInputSqlServerSqlDbSyncTaskProperties: %+v", err)
	}

	return encoded, nil
}
