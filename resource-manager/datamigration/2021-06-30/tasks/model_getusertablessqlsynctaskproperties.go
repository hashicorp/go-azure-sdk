package tasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProjectTaskProperties = GetUserTablesSqlSyncTaskProperties{}

type GetUserTablesSqlSyncTaskProperties struct {
	Input  *GetUserTablesSqlSyncTaskInput    `json:"input,omitempty"`
	Output *[]GetUserTablesSqlSyncTaskOutput `json:"output,omitempty"`

	// Fields inherited from ProjectTaskProperties
	ClientData *map[string]string   `json:"clientData,omitempty"`
	Commands   *[]CommandProperties `json:"commands,omitempty"`
	Errors     *[]ODataError        `json:"errors,omitempty"`
	State      *TaskState           `json:"state,omitempty"`
}

var _ json.Marshaler = GetUserTablesSqlSyncTaskProperties{}

func (s GetUserTablesSqlSyncTaskProperties) MarshalJSON() ([]byte, error) {
	type wrapper GetUserTablesSqlSyncTaskProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GetUserTablesSqlSyncTaskProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GetUserTablesSqlSyncTaskProperties: %+v", err)
	}
	decoded["taskType"] = "GetUserTables.AzureSqlDb.Sync"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GetUserTablesSqlSyncTaskProperties: %+v", err)
	}

	return encoded, nil
}
