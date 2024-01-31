package tasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProjectTaskProperties = MigrateSsisTaskProperties{}

type MigrateSsisTaskProperties struct {
	Input  *MigrateSsisTaskInput    `json:"input,omitempty"`
	Output *[]MigrateSsisTaskOutput `json:"output,omitempty"`

	// Fields inherited from ProjectTaskProperties
	ClientData *map[string]string   `json:"clientData,omitempty"`
	Commands   *[]CommandProperties `json:"commands,omitempty"`
	Errors     *[]ODataError        `json:"errors,omitempty"`
	State      *TaskState           `json:"state,omitempty"`
}

var _ json.Marshaler = MigrateSsisTaskProperties{}

func (s MigrateSsisTaskProperties) MarshalJSON() ([]byte, error) {
	type wrapper MigrateSsisTaskProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MigrateSsisTaskProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MigrateSsisTaskProperties: %+v", err)
	}
	decoded["taskType"] = "Migrate.Ssis"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MigrateSsisTaskProperties: %+v", err)
	}

	return encoded, nil
}
