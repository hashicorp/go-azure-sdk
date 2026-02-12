package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ JobData = CommandJobData{}

type CommandJobData struct {
	Path   string       `json:"path"`
	Target string       `json:"target"`
	Value  *interface{} `json:"value,omitempty"`

	// Fields inherited from JobData

	Type string `json:"type"`
}

func (s CommandJobData) JobData() BaseJobDataImpl {
	return BaseJobDataImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = CommandJobData{}

func (s CommandJobData) MarshalJSON() ([]byte, error) {
	type wrapper CommandJobData
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CommandJobData: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CommandJobData: %+v", err)
	}

	decoded["type"] = "command"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CommandJobData: %+v", err)
	}

	return encoded, nil
}
