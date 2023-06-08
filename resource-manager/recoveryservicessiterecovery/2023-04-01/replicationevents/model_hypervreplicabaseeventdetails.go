package replicationevents

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EventProviderSpecificDetails = HyperVReplicaBaseEventDetails{}

type HyperVReplicaBaseEventDetails struct {
	ContainerName       *string `json:"containerName,omitempty"`
	FabricName          *string `json:"fabricName,omitempty"`
	RemoteContainerName *string `json:"remoteContainerName,omitempty"`
	RemoteFabricName    *string `json:"remoteFabricName,omitempty"`

	// Fields inherited from EventProviderSpecificDetails
}

var _ json.Marshaler = HyperVReplicaBaseEventDetails{}

func (s HyperVReplicaBaseEventDetails) MarshalJSON() ([]byte, error) {
	type wrapper HyperVReplicaBaseEventDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling HyperVReplicaBaseEventDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling HyperVReplicaBaseEventDetails: %+v", err)
	}
	decoded["instanceType"] = "HyperVReplicaBaseEventDetails"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling HyperVReplicaBaseEventDetails: %+v", err)
	}

	return encoded, nil
}
