package tasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProjectTaskProperties = ConnectToTargetSqlMISyncTaskProperties{}

type ConnectToTargetSqlMISyncTaskProperties struct {
	Input  *ConnectToTargetSqlMISyncTaskInput    `json:"input,omitempty"`
	Output *[]ConnectToTargetSqlMISyncTaskOutput `json:"output,omitempty"`

	// Fields inherited from ProjectTaskProperties
	Commands *[]CommandProperties `json:"commands,omitempty"`
	Errors   *[]ODataError        `json:"errors,omitempty"`
	State    *TaskState           `json:"state,omitempty"`
}

var _ json.Marshaler = ConnectToTargetSqlMISyncTaskProperties{}

func (s ConnectToTargetSqlMISyncTaskProperties) MarshalJSON() ([]byte, error) {
	type wrapper ConnectToTargetSqlMISyncTaskProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ConnectToTargetSqlMISyncTaskProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ConnectToTargetSqlMISyncTaskProperties: %+v", err)
	}
	decoded["taskType"] = "ConnectToTarget.AzureSqlDbMI.Sync.LRS"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ConnectToTargetSqlMISyncTaskProperties: %+v", err)
	}

	return encoded, nil
}
