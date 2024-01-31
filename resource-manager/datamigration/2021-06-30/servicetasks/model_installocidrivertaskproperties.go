package servicetasks

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ProjectTaskProperties = InstallOCIDriverTaskProperties{}

type InstallOCIDriverTaskProperties struct {
	Input  *InstallOCIDriverTaskInput    `json:"input,omitempty"`
	Output *[]InstallOCIDriverTaskOutput `json:"output,omitempty"`

	// Fields inherited from ProjectTaskProperties
	ClientData *map[string]string   `json:"clientData,omitempty"`
	Commands   *[]CommandProperties `json:"commands,omitempty"`
	Errors     *[]ODataError        `json:"errors,omitempty"`
	State      *TaskState           `json:"state,omitempty"`
}

var _ json.Marshaler = InstallOCIDriverTaskProperties{}

func (s InstallOCIDriverTaskProperties) MarshalJSON() ([]byte, error) {
	type wrapper InstallOCIDriverTaskProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling InstallOCIDriverTaskProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling InstallOCIDriverTaskProperties: %+v", err)
	}
	decoded["taskType"] = "Service.Install.OCI"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling InstallOCIDriverTaskProperties: %+v", err)
	}

	return encoded, nil
}
