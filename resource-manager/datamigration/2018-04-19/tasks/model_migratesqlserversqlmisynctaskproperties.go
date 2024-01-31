package tasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProjectTaskProperties = MigrateSqlServerSqlMISyncTaskProperties{}

type MigrateSqlServerSqlMISyncTaskProperties struct {
	Input  *SqlServerSqlMISyncTaskInput           `json:"input,omitempty"`
	Output *[]MigrateSqlServerSqlMISyncTaskOutput `json:"output,omitempty"`

	// Fields inherited from ProjectTaskProperties
	Commands *[]CommandProperties `json:"commands,omitempty"`
	Errors   *[]ODataError        `json:"errors,omitempty"`
	State    *TaskState           `json:"state,omitempty"`
}

var _ json.Marshaler = MigrateSqlServerSqlMISyncTaskProperties{}

func (s MigrateSqlServerSqlMISyncTaskProperties) MarshalJSON() ([]byte, error) {
	type wrapper MigrateSqlServerSqlMISyncTaskProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MigrateSqlServerSqlMISyncTaskProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MigrateSqlServerSqlMISyncTaskProperties: %+v", err)
	}
	decoded["taskType"] = "Migrate.SqlServer.AzureSqlDbMI.Sync.LRS"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MigrateSqlServerSqlMISyncTaskProperties: %+v", err)
	}

	return encoded, nil
}
