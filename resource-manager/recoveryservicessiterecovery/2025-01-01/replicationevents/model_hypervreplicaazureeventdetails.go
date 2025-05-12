package replicationevents

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EventProviderSpecificDetails = HyperVReplicaAzureEventDetails{}

type HyperVReplicaAzureEventDetails struct {
	ContainerName       *string `json:"containerName,omitempty"`
	FabricName          *string `json:"fabricName,omitempty"`
	RemoteContainerName *string `json:"remoteContainerName,omitempty"`

	// Fields inherited from EventProviderSpecificDetails

	InstanceType string `json:"instanceType"`
}

func (s HyperVReplicaAzureEventDetails) EventProviderSpecificDetails() BaseEventProviderSpecificDetailsImpl {
	return BaseEventProviderSpecificDetailsImpl{
		InstanceType: s.InstanceType,
	}
}

var _ json.Marshaler = HyperVReplicaAzureEventDetails{}

func (s HyperVReplicaAzureEventDetails) MarshalJSON() ([]byte, error) {
	type wrapper HyperVReplicaAzureEventDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling HyperVReplicaAzureEventDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling HyperVReplicaAzureEventDetails: %+v", err)
	}

	decoded["instanceType"] = "HyperVReplicaAzure"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling HyperVReplicaAzureEventDetails: %+v", err)
	}

	return encoded, nil
}
