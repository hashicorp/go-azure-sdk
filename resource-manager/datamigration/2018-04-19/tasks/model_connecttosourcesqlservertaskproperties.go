package tasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProjectTaskProperties = ConnectToSourceSqlServerTaskProperties{}

type ConnectToSourceSqlServerTaskProperties struct {
	Input  *ConnectToSourceSqlServerTaskInput    `json:"input,omitempty"`
	Output *[]ConnectToSourceSqlServerTaskOutput `json:"output,omitempty"`

	// Fields inherited from ProjectTaskProperties
	Commands *[]CommandProperties `json:"commands,omitempty"`
	Errors   *[]ODataError        `json:"errors,omitempty"`
	State    *TaskState           `json:"state,omitempty"`
}

var _ json.Marshaler = ConnectToSourceSqlServerTaskProperties{}

func (s ConnectToSourceSqlServerTaskProperties) MarshalJSON() ([]byte, error) {
	type wrapper ConnectToSourceSqlServerTaskProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ConnectToSourceSqlServerTaskProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ConnectToSourceSqlServerTaskProperties: %+v", err)
	}
	decoded["taskType"] = "ConnectToSource.SqlServer"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ConnectToSourceSqlServerTaskProperties: %+v", err)
	}

	return encoded, nil
}
