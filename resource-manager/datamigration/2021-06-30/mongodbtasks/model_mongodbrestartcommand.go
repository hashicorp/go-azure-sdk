package mongodbtasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CommandProperties = MongoDbRestartCommand{}

type MongoDbRestartCommand struct {
	Input *MongoDbCommandInput `json:"input,omitempty"`

	// Fields inherited from CommandProperties
	Errors *[]ODataError `json:"errors,omitempty"`
	State  *CommandState `json:"state,omitempty"`
}

var _ json.Marshaler = MongoDbRestartCommand{}

func (s MongoDbRestartCommand) MarshalJSON() ([]byte, error) {
	type wrapper MongoDbRestartCommand
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MongoDbRestartCommand: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MongoDbRestartCommand: %+v", err)
	}
	decoded["commandType"] = "restart"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MongoDbRestartCommand: %+v", err)
	}

	return encoded, nil
}
