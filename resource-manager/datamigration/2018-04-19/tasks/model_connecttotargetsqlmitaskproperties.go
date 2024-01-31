package tasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProjectTaskProperties = ConnectToTargetSqlMITaskProperties{}

type ConnectToTargetSqlMITaskProperties struct {
	Input  *ConnectToTargetSqlMITaskInput    `json:"input,omitempty"`
	Output *[]ConnectToTargetSqlMITaskOutput `json:"output,omitempty"`

	// Fields inherited from ProjectTaskProperties
	Commands *[]CommandProperties `json:"commands,omitempty"`
	Errors   *[]ODataError        `json:"errors,omitempty"`
	State    *TaskState           `json:"state,omitempty"`
}

var _ json.Marshaler = ConnectToTargetSqlMITaskProperties{}

func (s ConnectToTargetSqlMITaskProperties) MarshalJSON() ([]byte, error) {
	type wrapper ConnectToTargetSqlMITaskProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ConnectToTargetSqlMITaskProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ConnectToTargetSqlMITaskProperties: %+v", err)
	}
	decoded["taskType"] = "ConnectToTarget.AzureSqlDbMI"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ConnectToTargetSqlMITaskProperties: %+v", err)
	}

	return encoded, nil
}
