package tasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProjectTaskProperties = GetUserTablesSqlTaskProperties{}

type GetUserTablesSqlTaskProperties struct {
	Input  *GetUserTablesSqlTaskInput    `json:"input,omitempty"`
	Output *[]GetUserTablesSqlTaskOutput `json:"output,omitempty"`

	// Fields inherited from ProjectTaskProperties
	Commands *[]CommandProperties `json:"commands,omitempty"`
	Errors   *[]ODataError        `json:"errors,omitempty"`
	State    *TaskState           `json:"state,omitempty"`
}

var _ json.Marshaler = GetUserTablesSqlTaskProperties{}

func (s GetUserTablesSqlTaskProperties) MarshalJSON() ([]byte, error) {
	type wrapper GetUserTablesSqlTaskProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GetUserTablesSqlTaskProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GetUserTablesSqlTaskProperties: %+v", err)
	}
	decoded["taskType"] = "GetUserTables.Sql"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GetUserTablesSqlTaskProperties: %+v", err)
	}

	return encoded, nil
}
