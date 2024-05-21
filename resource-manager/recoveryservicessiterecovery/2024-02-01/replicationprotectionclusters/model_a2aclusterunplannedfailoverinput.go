package replicationprotectionclusters

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ClusterUnplannedFailoverProviderSpecificInput = A2AClusterUnplannedFailoverInput{}

type A2AClusterUnplannedFailoverInput struct {
	ClusterRecoveryPointId       *string   `json:"clusterRecoveryPointId,omitempty"`
	IndividualNodeRecoveryPoints *[]string `json:"individualNodeRecoveryPoints,omitempty"`

	// Fields inherited from ClusterUnplannedFailoverProviderSpecificInput
}

var _ json.Marshaler = A2AClusterUnplannedFailoverInput{}

func (s A2AClusterUnplannedFailoverInput) MarshalJSON() ([]byte, error) {
	type wrapper A2AClusterUnplannedFailoverInput
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling A2AClusterUnplannedFailoverInput: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling A2AClusterUnplannedFailoverInput: %+v", err)
	}
	decoded["instanceType"] = "A2A"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling A2AClusterUnplannedFailoverInput: %+v", err)
	}

	return encoded, nil
}
