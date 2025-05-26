package replicationjobs

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TaskTypeDetails = ScriptActionTaskDetails{}

type ScriptActionTaskDetails struct {
	IsPrimarySideScript *bool   `json:"isPrimarySideScript,omitempty"`
	Name                *string `json:"name,omitempty"`
	Output              *string `json:"output,omitempty"`
	Path                *string `json:"path,omitempty"`

	// Fields inherited from TaskTypeDetails

	InstanceType string `json:"instanceType"`
}

func (s ScriptActionTaskDetails) TaskTypeDetails() BaseTaskTypeDetailsImpl {
	return BaseTaskTypeDetailsImpl{
		InstanceType: s.InstanceType,
	}
}

var _ json.Marshaler = ScriptActionTaskDetails{}

func (s ScriptActionTaskDetails) MarshalJSON() ([]byte, error) {
	type wrapper ScriptActionTaskDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ScriptActionTaskDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ScriptActionTaskDetails: %+v", err)
	}

	decoded["instanceType"] = "ScriptActionTaskDetails"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ScriptActionTaskDetails: %+v", err)
	}

	return encoded, nil
}
