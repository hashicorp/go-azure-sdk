package tasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProjectTaskProperties = ConnectToSourcePostgreSqlSyncTaskProperties{}

type ConnectToSourcePostgreSqlSyncTaskProperties struct {
	Input  *ConnectToSourcePostgreSqlSyncTaskInput    `json:"input,omitempty"`
	Output *[]ConnectToSourcePostgreSqlSyncTaskOutput `json:"output,omitempty"`

	// Fields inherited from ProjectTaskProperties
	ClientData *map[string]string   `json:"clientData,omitempty"`
	Commands   *[]CommandProperties `json:"commands,omitempty"`
	Errors     *[]ODataError        `json:"errors,omitempty"`
	State      *TaskState           `json:"state,omitempty"`
}

var _ json.Marshaler = ConnectToSourcePostgreSqlSyncTaskProperties{}

func (s ConnectToSourcePostgreSqlSyncTaskProperties) MarshalJSON() ([]byte, error) {
	type wrapper ConnectToSourcePostgreSqlSyncTaskProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ConnectToSourcePostgreSqlSyncTaskProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ConnectToSourcePostgreSqlSyncTaskProperties: %+v", err)
	}
	decoded["taskType"] = "ConnectToSource.PostgreSql.Sync"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ConnectToSourcePostgreSqlSyncTaskProperties: %+v", err)
	}

	return encoded, nil
}
