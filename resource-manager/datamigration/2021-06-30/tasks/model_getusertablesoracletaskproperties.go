package tasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProjectTaskProperties = GetUserTablesOracleTaskProperties{}

type GetUserTablesOracleTaskProperties struct {
	Input  *GetUserTablesOracleTaskInput    `json:"input,omitempty"`
	Output *[]GetUserTablesOracleTaskOutput `json:"output,omitempty"`

	// Fields inherited from ProjectTaskProperties
	ClientData *map[string]string   `json:"clientData,omitempty"`
	Commands   *[]CommandProperties `json:"commands,omitempty"`
	Errors     *[]ODataError        `json:"errors,omitempty"`
	State      *TaskState           `json:"state,omitempty"`
}

var _ json.Marshaler = GetUserTablesOracleTaskProperties{}

func (s GetUserTablesOracleTaskProperties) MarshalJSON() ([]byte, error) {
	type wrapper GetUserTablesOracleTaskProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GetUserTablesOracleTaskProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GetUserTablesOracleTaskProperties: %+v", err)
	}
	decoded["taskType"] = "GetUserTablesOracle"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GetUserTablesOracleTaskProperties: %+v", err)
	}

	return encoded, nil
}
