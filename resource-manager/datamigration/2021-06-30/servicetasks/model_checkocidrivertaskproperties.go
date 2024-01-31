package servicetasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProjectTaskProperties = CheckOCIDriverTaskProperties{}

type CheckOCIDriverTaskProperties struct {
	Input  *CheckOCIDriverTaskInput    `json:"input,omitempty"`
	Output *[]CheckOCIDriverTaskOutput `json:"output,omitempty"`

	// Fields inherited from ProjectTaskProperties
	ClientData *map[string]string   `json:"clientData,omitempty"`
	Commands   *[]CommandProperties `json:"commands,omitempty"`
	Errors     *[]ODataError        `json:"errors,omitempty"`
	State      *TaskState           `json:"state,omitempty"`
}

var _ json.Marshaler = CheckOCIDriverTaskProperties{}

func (s CheckOCIDriverTaskProperties) MarshalJSON() ([]byte, error) {
	type wrapper CheckOCIDriverTaskProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CheckOCIDriverTaskProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CheckOCIDriverTaskProperties: %+v", err)
	}
	decoded["taskType"] = "Service.Check.OCI"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CheckOCIDriverTaskProperties: %+v", err)
	}

	return encoded, nil
}
