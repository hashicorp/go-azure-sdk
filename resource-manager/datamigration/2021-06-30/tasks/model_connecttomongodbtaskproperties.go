package tasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProjectTaskProperties = ConnectToMongoDbTaskProperties{}

type ConnectToMongoDbTaskProperties struct {
	Input  *MongoDbConnectionInfo `json:"input,omitempty"`
	Output *[]MongoDbClusterInfo  `json:"output,omitempty"`

	// Fields inherited from ProjectTaskProperties
	ClientData *map[string]string   `json:"clientData,omitempty"`
	Commands   *[]CommandProperties `json:"commands,omitempty"`
	Errors     *[]ODataError        `json:"errors,omitempty"`
	State      *TaskState           `json:"state,omitempty"`
}

var _ json.Marshaler = ConnectToMongoDbTaskProperties{}

func (s ConnectToMongoDbTaskProperties) MarshalJSON() ([]byte, error) {
	type wrapper ConnectToMongoDbTaskProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ConnectToMongoDbTaskProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ConnectToMongoDbTaskProperties: %+v", err)
	}
	decoded["taskType"] = "Connect.MongoDb"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ConnectToMongoDbTaskProperties: %+v", err)
	}

	return encoded, nil
}
