package replicationjobs

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TaskTypeDetails = ConsistencyCheckTaskDetails{}

type ConsistencyCheckTaskDetails struct {
	VmDetails *[]InconsistentVmDetails `json:"vmDetails,omitempty"`

	// Fields inherited from TaskTypeDetails
}

var _ json.Marshaler = ConsistencyCheckTaskDetails{}

func (s ConsistencyCheckTaskDetails) MarshalJSON() ([]byte, error) {
	type wrapper ConsistencyCheckTaskDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ConsistencyCheckTaskDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ConsistencyCheckTaskDetails: %+v", err)
	}
	decoded["instanceType"] = "ConsistencyCheckTaskDetails"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ConsistencyCheckTaskDetails: %+v", err)
	}

	return encoded, nil
}
