package replicationjobs

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ JobDetails = TestFailoverJobDetails{}

type TestFailoverJobDetails struct {
	Comments             *string                                    `json:"comments,omitempty"`
	NetworkFriendlyName  *string                                    `json:"networkFriendlyName,omitempty"`
	NetworkName          *string                                    `json:"networkName,omitempty"`
	NetworkType          *string                                    `json:"networkType,omitempty"`
	ProtectedItemDetails *[]FailoverReplicationProtectedItemDetails `json:"protectedItemDetails,omitempty"`
	TestFailoverStatus   *string                                    `json:"testFailoverStatus,omitempty"`

	// Fields inherited from JobDetails

	AffectedObjectDetails *map[string]string `json:"affectedObjectDetails,omitempty"`
	InstanceType          string             `json:"instanceType"`
}

func (s TestFailoverJobDetails) JobDetails() BaseJobDetailsImpl {
	return BaseJobDetailsImpl{
		AffectedObjectDetails: s.AffectedObjectDetails,
		InstanceType:          s.InstanceType,
	}
}

var _ json.Marshaler = TestFailoverJobDetails{}

func (s TestFailoverJobDetails) MarshalJSON() ([]byte, error) {
	type wrapper TestFailoverJobDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TestFailoverJobDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TestFailoverJobDetails: %+v", err)
	}

	decoded["instanceType"] = "TestFailoverJobDetails"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TestFailoverJobDetails: %+v", err)
	}

	return encoded, nil
}
