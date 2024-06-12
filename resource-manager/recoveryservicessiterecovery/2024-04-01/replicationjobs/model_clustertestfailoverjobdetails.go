package replicationjobs

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ JobDetails = ClusterTestFailoverJobDetails{}

type ClusterTestFailoverJobDetails struct {
	Comments             *string                                    `json:"comments,omitempty"`
	NetworkFriendlyName  *string                                    `json:"networkFriendlyName,omitempty"`
	NetworkName          *string                                    `json:"networkName,omitempty"`
	NetworkType          *string                                    `json:"networkType,omitempty"`
	ProtectedItemDetails *[]FailoverReplicationProtectedItemDetails `json:"protectedItemDetails,omitempty"`
	TestFailoverStatus   *string                                    `json:"testFailoverStatus,omitempty"`

	// Fields inherited from JobDetails
	AffectedObjectDetails *map[string]string `json:"affectedObjectDetails,omitempty"`
}

var _ json.Marshaler = ClusterTestFailoverJobDetails{}

func (s ClusterTestFailoverJobDetails) MarshalJSON() ([]byte, error) {
	type wrapper ClusterTestFailoverJobDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ClusterTestFailoverJobDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ClusterTestFailoverJobDetails: %+v", err)
	}
	decoded["instanceType"] = "ClusterTestFailoverJobDetails"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ClusterTestFailoverJobDetails: %+v", err)
	}

	return encoded, nil
}
