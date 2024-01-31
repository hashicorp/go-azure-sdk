package tasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProjectTaskProperties = ConnectToTargetSqlDbTaskProperties{}

type ConnectToTargetSqlDbTaskProperties struct {
	Input  *ConnectToTargetSqlDbTaskInput    `json:"input,omitempty"`
	Output *[]ConnectToTargetSqlDbTaskOutput `json:"output,omitempty"`

	// Fields inherited from ProjectTaskProperties
	Commands *[]CommandProperties `json:"commands,omitempty"`
	Errors   *[]ODataError        `json:"errors,omitempty"`
	State    *TaskState           `json:"state,omitempty"`
}

var _ json.Marshaler = ConnectToTargetSqlDbTaskProperties{}

func (s ConnectToTargetSqlDbTaskProperties) MarshalJSON() ([]byte, error) {
	type wrapper ConnectToTargetSqlDbTaskProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ConnectToTargetSqlDbTaskProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ConnectToTargetSqlDbTaskProperties: %+v", err)
	}
	decoded["taskType"] = "ConnectToTarget.SqlDb"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ConnectToTargetSqlDbTaskProperties: %+v", err)
	}

	return encoded, nil
}
