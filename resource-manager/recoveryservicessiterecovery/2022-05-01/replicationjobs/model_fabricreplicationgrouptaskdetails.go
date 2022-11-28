package replicationjobs

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TaskTypeDetails = FabricReplicationGroupTaskDetails{}

type FabricReplicationGroupTaskDetails struct {
	SkippedReason       *string `json:"skippedReason,omitempty"`
	SkippedReasonString *string `json:"skippedReasonString,omitempty"`

	// Fields inherited from JobTaskDetails
	JobTask *JobEntity `json:"jobTask"`
}

var _ json.Marshaler = FabricReplicationGroupTaskDetails{}

func (s FabricReplicationGroupTaskDetails) MarshalJSON() ([]byte, error) {
	type wrapper FabricReplicationGroupTaskDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling FabricReplicationGroupTaskDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling FabricReplicationGroupTaskDetails: %+v", err)
	}
	decoded["instanceType"] = "FabricReplicationGroupTaskDetails"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling FabricReplicationGroupTaskDetails: %+v", err)
	}

	return encoded, nil
}
