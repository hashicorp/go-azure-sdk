package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ JobData = CloudPropertyJobData{}

type CloudPropertyJobData struct {
	Path   string       `json:"path"`
	Target string       `json:"target"`
	Value  *interface{} `json:"value,omitempty"`

	// Fields inherited from JobData

	Type string `json:"type"`
}

func (s CloudPropertyJobData) JobData() BaseJobDataImpl {
	return BaseJobDataImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = CloudPropertyJobData{}

func (s CloudPropertyJobData) MarshalJSON() ([]byte, error) {
	type wrapper CloudPropertyJobData
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CloudPropertyJobData: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CloudPropertyJobData: %+v", err)
	}

	decoded["type"] = "cloudProperty"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CloudPropertyJobData: %+v", err)
	}

	return encoded, nil
}
