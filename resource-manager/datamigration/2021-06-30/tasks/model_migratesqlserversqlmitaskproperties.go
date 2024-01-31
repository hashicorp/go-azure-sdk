package tasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProjectTaskProperties = MigrateSqlServerSqlMITaskProperties{}

type MigrateSqlServerSqlMITaskProperties struct {
	Input  *MigrateSqlServerSqlMITaskInput    `json:"input,omitempty"`
	Output *[]MigrateSqlServerSqlMITaskOutput `json:"output,omitempty"`

	// Fields inherited from ProjectTaskProperties
	ClientData *map[string]string   `json:"clientData,omitempty"`
	Commands   *[]CommandProperties `json:"commands,omitempty"`
	Errors     *[]ODataError        `json:"errors,omitempty"`
	State      *TaskState           `json:"state,omitempty"`
}

var _ json.Marshaler = MigrateSqlServerSqlMITaskProperties{}

func (s MigrateSqlServerSqlMITaskProperties) MarshalJSON() ([]byte, error) {
	type wrapper MigrateSqlServerSqlMITaskProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MigrateSqlServerSqlMITaskProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MigrateSqlServerSqlMITaskProperties: %+v", err)
	}
	decoded["taskType"] = "Migrate.SqlServer.AzureSqlDbMI"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MigrateSqlServerSqlMITaskProperties: %+v", err)
	}

	return encoded, nil
}
