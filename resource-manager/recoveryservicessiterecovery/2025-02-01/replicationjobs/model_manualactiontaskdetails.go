package replicationjobs

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TaskTypeDetails = ManualActionTaskDetails{}

type ManualActionTaskDetails struct {
	Instructions *string `json:"instructions,omitempty"`
	Name         *string `json:"name,omitempty"`
	Observation  *string `json:"observation,omitempty"`

	// Fields inherited from TaskTypeDetails

	InstanceType string `json:"instanceType"`
}

func (s ManualActionTaskDetails) TaskTypeDetails() BaseTaskTypeDetailsImpl {
	return BaseTaskTypeDetailsImpl{
		InstanceType: s.InstanceType,
	}
}

var _ json.Marshaler = ManualActionTaskDetails{}

func (s ManualActionTaskDetails) MarshalJSON() ([]byte, error) {
	type wrapper ManualActionTaskDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ManualActionTaskDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ManualActionTaskDetails: %+v", err)
	}

	decoded["instanceType"] = "ManualActionTaskDetails"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ManualActionTaskDetails: %+v", err)
	}

	return encoded, nil
}
