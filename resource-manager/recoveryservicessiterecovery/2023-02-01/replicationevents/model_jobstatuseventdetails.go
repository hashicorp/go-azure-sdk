package replicationevents

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EventSpecificDetails = JobStatusEventDetails{}

type JobStatusEventDetails struct {
	AffectedObjectType *string `json:"affectedObjectType,omitempty"`
	JobFriendlyName    *string `json:"jobFriendlyName,omitempty"`
	JobId              *string `json:"jobId,omitempty"`
	JobStatus          *string `json:"jobStatus,omitempty"`

	// Fields inherited from EventSpecificDetails
}

var _ json.Marshaler = JobStatusEventDetails{}

func (s JobStatusEventDetails) MarshalJSON() ([]byte, error) {
	type wrapper JobStatusEventDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling JobStatusEventDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling JobStatusEventDetails: %+v", err)
	}
	decoded["instanceType"] = "JobStatus"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling JobStatusEventDetails: %+v", err)
	}

	return encoded, nil
}
