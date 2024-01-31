package tasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProjectTaskProperties = MigrateSqlServerSqlDbSyncTaskProperties{}

type MigrateSqlServerSqlDbSyncTaskProperties struct {
	Input  *MigrateSqlServerSqlDbSyncTaskInput    `json:"input,omitempty"`
	Output *[]MigrateSqlServerSqlDbSyncTaskOutput `json:"output,omitempty"`

	// Fields inherited from ProjectTaskProperties
	Commands *[]CommandProperties `json:"commands,omitempty"`
	Errors   *[]ODataError        `json:"errors,omitempty"`
	State    *TaskState           `json:"state,omitempty"`
}

var _ json.Marshaler = MigrateSqlServerSqlDbSyncTaskProperties{}

func (s MigrateSqlServerSqlDbSyncTaskProperties) MarshalJSON() ([]byte, error) {
	type wrapper MigrateSqlServerSqlDbSyncTaskProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MigrateSqlServerSqlDbSyncTaskProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MigrateSqlServerSqlDbSyncTaskProperties: %+v", err)
	}
	decoded["taskType"] = "Migrate.SqlServer.AzureSqlDb.Sync"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MigrateSqlServerSqlDbSyncTaskProperties: %+v", err)
	}

	return encoded, nil
}
