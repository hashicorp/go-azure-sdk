package replicationprotectionclusters

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ClusterTestFailoverProviderSpecificInput = A2AClusterTestFailoverInput{}

type A2AClusterTestFailoverInput struct {
	ClusterRecoveryPointId       *string   `json:"clusterRecoveryPointId,omitempty"`
	IndividualNodeRecoveryPoints *[]string `json:"individualNodeRecoveryPoints,omitempty"`

	// Fields inherited from ClusterTestFailoverProviderSpecificInput

	InstanceType string `json:"instanceType"`
}

func (s A2AClusterTestFailoverInput) ClusterTestFailoverProviderSpecificInput() BaseClusterTestFailoverProviderSpecificInputImpl {
	return BaseClusterTestFailoverProviderSpecificInputImpl{
		InstanceType: s.InstanceType,
	}
}

var _ json.Marshaler = A2AClusterTestFailoverInput{}

func (s A2AClusterTestFailoverInput) MarshalJSON() ([]byte, error) {
	type wrapper A2AClusterTestFailoverInput
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling A2AClusterTestFailoverInput: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling A2AClusterTestFailoverInput: %+v", err)
	}

	decoded["instanceType"] = "A2A"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling A2AClusterTestFailoverInput: %+v", err)
	}

	return encoded, nil
}
