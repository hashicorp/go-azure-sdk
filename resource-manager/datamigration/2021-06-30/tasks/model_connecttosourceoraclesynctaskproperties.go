package tasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProjectTaskProperties = ConnectToSourceOracleSyncTaskProperties{}

type ConnectToSourceOracleSyncTaskProperties struct {
	Input  *ConnectToSourceOracleSyncTaskInput    `json:"input,omitempty"`
	Output *[]ConnectToSourceOracleSyncTaskOutput `json:"output,omitempty"`

	// Fields inherited from ProjectTaskProperties
	ClientData *map[string]string   `json:"clientData,omitempty"`
	Commands   *[]CommandProperties `json:"commands,omitempty"`
	Errors     *[]ODataError        `json:"errors,omitempty"`
	State      *TaskState           `json:"state,omitempty"`
}

var _ json.Marshaler = ConnectToSourceOracleSyncTaskProperties{}

func (s ConnectToSourceOracleSyncTaskProperties) MarshalJSON() ([]byte, error) {
	type wrapper ConnectToSourceOracleSyncTaskProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ConnectToSourceOracleSyncTaskProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ConnectToSourceOracleSyncTaskProperties: %+v", err)
	}
	decoded["taskType"] = "ConnectToSource.Oracle.Sync"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ConnectToSourceOracleSyncTaskProperties: %+v", err)
	}

	return encoded, nil
}
