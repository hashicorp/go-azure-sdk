package replicationjobs

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TaskTypeDetails = JobTaskDetails{}

type JobTaskDetails struct {
	JobTask *JobEntity `json:"jobTask,omitempty"`

	// Fields inherited from TaskTypeDetails
}

var _ json.Marshaler = JobTaskDetails{}

func (s JobTaskDetails) MarshalJSON() ([]byte, error) {
	type wrapper JobTaskDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling JobTaskDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling JobTaskDetails: %+v", err)
	}
	decoded["instanceType"] = "JobTaskDetails"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling JobTaskDetails: %+v", err)
	}

	return encoded, nil
}
