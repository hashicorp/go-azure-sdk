package tasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProjectTaskProperties = MigratePostgreSqlAzureDbForPostgreSqlSyncTaskProperties{}

type MigratePostgreSqlAzureDbForPostgreSqlSyncTaskProperties struct {
	Input  *MigratePostgreSqlAzureDbForPostgreSqlSyncTaskInput    `json:"input,omitempty"`
	Output *[]MigratePostgreSqlAzureDbForPostgreSqlSyncTaskOutput `json:"output,omitempty"`

	// Fields inherited from ProjectTaskProperties
	Commands *[]CommandProperties `json:"commands,omitempty"`
	Errors   *[]ODataError        `json:"errors,omitempty"`
	State    *TaskState           `json:"state,omitempty"`
}

var _ json.Marshaler = MigratePostgreSqlAzureDbForPostgreSqlSyncTaskProperties{}

func (s MigratePostgreSqlAzureDbForPostgreSqlSyncTaskProperties) MarshalJSON() ([]byte, error) {
	type wrapper MigratePostgreSqlAzureDbForPostgreSqlSyncTaskProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MigratePostgreSqlAzureDbForPostgreSqlSyncTaskProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MigratePostgreSqlAzureDbForPostgreSqlSyncTaskProperties: %+v", err)
	}
	decoded["taskType"] = "Migrate.PostgreSql.AzureDbForPostgreSql.Sync"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MigratePostgreSqlAzureDbForPostgreSqlSyncTaskProperties: %+v", err)
	}

	return encoded, nil
}
