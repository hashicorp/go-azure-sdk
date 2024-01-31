package tasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProjectTaskProperties = ValidateMongoDbTaskProperties{}

type ValidateMongoDbTaskProperties struct {
	Input  *MongoDbMigrationSettings   `json:"input,omitempty"`
	Output *[]MongoDbMigrationProgress `json:"output,omitempty"`

	// Fields inherited from ProjectTaskProperties
	ClientData *map[string]string   `json:"clientData,omitempty"`
	Commands   *[]CommandProperties `json:"commands,omitempty"`
	Errors     *[]ODataError        `json:"errors,omitempty"`
	State      *TaskState           `json:"state,omitempty"`
}

var _ json.Marshaler = ValidateMongoDbTaskProperties{}

func (s ValidateMongoDbTaskProperties) MarshalJSON() ([]byte, error) {
	type wrapper ValidateMongoDbTaskProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ValidateMongoDbTaskProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ValidateMongoDbTaskProperties: %+v", err)
	}
	decoded["taskType"] = "Validate.MongoDb"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ValidateMongoDbTaskProperties: %+v", err)
	}

	return encoded, nil
}
