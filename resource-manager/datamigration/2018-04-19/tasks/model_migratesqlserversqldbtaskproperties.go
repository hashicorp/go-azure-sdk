package tasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProjectTaskProperties = MigrateSqlServerSqlDbTaskProperties{}

type MigrateSqlServerSqlDbTaskProperties struct {
	Input  *MigrateSqlServerSqlDbTaskInput    `json:"input,omitempty"`
	Output *[]MigrateSqlServerSqlDbTaskOutput `json:"output,omitempty"`

	// Fields inherited from ProjectTaskProperties
	Commands *[]CommandProperties `json:"commands,omitempty"`
	Errors   *[]ODataError        `json:"errors,omitempty"`
	State    *TaskState           `json:"state,omitempty"`
}

var _ json.Marshaler = MigrateSqlServerSqlDbTaskProperties{}

func (s MigrateSqlServerSqlDbTaskProperties) MarshalJSON() ([]byte, error) {
	type wrapper MigrateSqlServerSqlDbTaskProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MigrateSqlServerSqlDbTaskProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MigrateSqlServerSqlDbTaskProperties: %+v", err)
	}
	decoded["taskType"] = "Migrate.SqlServer.SqlDb"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MigrateSqlServerSqlDbTaskProperties: %+v", err)
	}

	return encoded, nil
}
