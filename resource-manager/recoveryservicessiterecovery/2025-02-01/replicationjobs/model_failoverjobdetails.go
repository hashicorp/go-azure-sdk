package replicationjobs

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ JobDetails = FailoverJobDetails{}

type FailoverJobDetails struct {
	ProtectedItemDetails *[]FailoverReplicationProtectedItemDetails `json:"protectedItemDetails,omitempty"`

	// Fields inherited from JobDetails

	AffectedObjectDetails *map[string]string `json:"affectedObjectDetails,omitempty"`
	InstanceType          string             `json:"instanceType"`
}

func (s FailoverJobDetails) JobDetails() BaseJobDetailsImpl {
	return BaseJobDetailsImpl{
		AffectedObjectDetails: s.AffectedObjectDetails,
		InstanceType:          s.InstanceType,
	}
}

var _ json.Marshaler = FailoverJobDetails{}

func (s FailoverJobDetails) MarshalJSON() ([]byte, error) {
	type wrapper FailoverJobDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling FailoverJobDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling FailoverJobDetails: %+v", err)
	}

	decoded["instanceType"] = "FailoverJobDetails"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling FailoverJobDetails: %+v", err)
	}

	return encoded, nil
}
