package tasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProjectTaskProperties = GetUserTablesMySqlTaskProperties{}

type GetUserTablesMySqlTaskProperties struct {
	Input  *GetUserTablesMySqlTaskInput    `json:"input,omitempty"`
	Output *[]GetUserTablesMySqlTaskOutput `json:"output,omitempty"`

	// Fields inherited from ProjectTaskProperties
	ClientData *map[string]string   `json:"clientData,omitempty"`
	Commands   *[]CommandProperties `json:"commands,omitempty"`
	Errors     *[]ODataError        `json:"errors,omitempty"`
	State      *TaskState           `json:"state,omitempty"`
}

var _ json.Marshaler = GetUserTablesMySqlTaskProperties{}

func (s GetUserTablesMySqlTaskProperties) MarshalJSON() ([]byte, error) {
	type wrapper GetUserTablesMySqlTaskProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GetUserTablesMySqlTaskProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GetUserTablesMySqlTaskProperties: %+v", err)
	}
	decoded["taskType"] = "GetUserTablesMySql"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GetUserTablesMySqlTaskProperties: %+v", err)
	}

	return encoded, nil
}
