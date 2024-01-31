package tasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProjectTaskProperties = ValidateMigrationInputSqlServerSqlMISyncTaskProperties{}

type ValidateMigrationInputSqlServerSqlMISyncTaskProperties struct {
	Input  *SqlServerSqlMISyncTaskInput                          `json:"input,omitempty"`
	Output *[]ValidateMigrationInputSqlServerSqlMISyncTaskOutput `json:"output,omitempty"`

	// Fields inherited from ProjectTaskProperties
	Commands *[]CommandProperties `json:"commands,omitempty"`
	Errors   *[]ODataError        `json:"errors,omitempty"`
	State    *TaskState           `json:"state,omitempty"`
}

var _ json.Marshaler = ValidateMigrationInputSqlServerSqlMISyncTaskProperties{}

func (s ValidateMigrationInputSqlServerSqlMISyncTaskProperties) MarshalJSON() ([]byte, error) {
	type wrapper ValidateMigrationInputSqlServerSqlMISyncTaskProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ValidateMigrationInputSqlServerSqlMISyncTaskProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ValidateMigrationInputSqlServerSqlMISyncTaskProperties: %+v", err)
	}
	decoded["taskType"] = "ValidateMigrationInput.SqlServer.AzureSqlDbMI.Sync.LRS"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ValidateMigrationInputSqlServerSqlMISyncTaskProperties: %+v", err)
	}

	return encoded, nil
}
