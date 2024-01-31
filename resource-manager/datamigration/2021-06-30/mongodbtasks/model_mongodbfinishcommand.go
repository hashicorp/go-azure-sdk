package mongodbtasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CommandProperties = MongoDbFinishCommand{}

type MongoDbFinishCommand struct {
	Input *MongoDbFinishCommandInput `json:"input,omitempty"`

	// Fields inherited from CommandProperties
	Errors *[]ODataError `json:"errors,omitempty"`
	State  *CommandState `json:"state,omitempty"`
}

var _ json.Marshaler = MongoDbFinishCommand{}

func (s MongoDbFinishCommand) MarshalJSON() ([]byte, error) {
	type wrapper MongoDbFinishCommand
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MongoDbFinishCommand: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MongoDbFinishCommand: %+v", err)
	}
	decoded["commandType"] = "finish"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MongoDbFinishCommand: %+v", err)
	}

	return encoded, nil
}
