package tasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProjectTaskProperties = ConnectToSourceSqlServerSyncTaskProperties{}

type ConnectToSourceSqlServerSyncTaskProperties struct {
	Input  *ConnectToSourceSqlServerTaskInput    `json:"input,omitempty"`
	Output *[]ConnectToSourceSqlServerTaskOutput `json:"output,omitempty"`

	// Fields inherited from ProjectTaskProperties
	Commands *[]CommandProperties `json:"commands,omitempty"`
	Errors   *[]ODataError        `json:"errors,omitempty"`
	State    *TaskState           `json:"state,omitempty"`
}

var _ json.Marshaler = ConnectToSourceSqlServerSyncTaskProperties{}

func (s ConnectToSourceSqlServerSyncTaskProperties) MarshalJSON() ([]byte, error) {
	type wrapper ConnectToSourceSqlServerSyncTaskProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ConnectToSourceSqlServerSyncTaskProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ConnectToSourceSqlServerSyncTaskProperties: %+v", err)
	}
	decoded["taskType"] = "ConnectToSource.SqlServer.Sync"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ConnectToSourceSqlServerSyncTaskProperties: %+v", err)
	}

	return encoded, nil
}
