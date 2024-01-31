package connecttosourcemysqltasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProjectTaskProperties = ConnectToSourceMySqlTaskProperties{}

type ConnectToSourceMySqlTaskProperties struct {
	Input  *ConnectToSourceMySqlTaskInput     `json:"input,omitempty"`
	Output *[]ConnectToSourceNonSqlTaskOutput `json:"output,omitempty"`

	// Fields inherited from ProjectTaskProperties
	Commands *[]CommandProperties `json:"commands,omitempty"`
	Errors   *[]ODataError        `json:"errors,omitempty"`
	State    *TaskState           `json:"state,omitempty"`
}

var _ json.Marshaler = ConnectToSourceMySqlTaskProperties{}

func (s ConnectToSourceMySqlTaskProperties) MarshalJSON() ([]byte, error) {
	type wrapper ConnectToSourceMySqlTaskProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ConnectToSourceMySqlTaskProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ConnectToSourceMySqlTaskProperties: %+v", err)
	}
	decoded["taskType"] = "ConnectToSource.MySql"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ConnectToSourceMySqlTaskProperties: %+v", err)
	}

	return encoded, nil
}
