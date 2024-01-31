package tasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProjectTaskProperties = ConnectToTargetAzureDbForMySqlTaskProperties{}

type ConnectToTargetAzureDbForMySqlTaskProperties struct {
	Input  *ConnectToTargetAzureDbForMySqlTaskInput    `json:"input,omitempty"`
	Output *[]ConnectToTargetAzureDbForMySqlTaskOutput `json:"output,omitempty"`

	// Fields inherited from ProjectTaskProperties
	Commands *[]CommandProperties `json:"commands,omitempty"`
	Errors   *[]ODataError        `json:"errors,omitempty"`
	State    *TaskState           `json:"state,omitempty"`
}

var _ json.Marshaler = ConnectToTargetAzureDbForMySqlTaskProperties{}

func (s ConnectToTargetAzureDbForMySqlTaskProperties) MarshalJSON() ([]byte, error) {
	type wrapper ConnectToTargetAzureDbForMySqlTaskProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ConnectToTargetAzureDbForMySqlTaskProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ConnectToTargetAzureDbForMySqlTaskProperties: %+v", err)
	}
	decoded["taskType"] = "ConnectToTarget.AzureDbForMySql"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ConnectToTargetAzureDbForMySqlTaskProperties: %+v", err)
	}

	return encoded, nil
}
