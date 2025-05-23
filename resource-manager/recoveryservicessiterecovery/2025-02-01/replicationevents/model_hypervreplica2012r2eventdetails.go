package replicationevents

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ EventProviderSpecificDetails = HyperVReplica2012R2EventDetails{}

type HyperVReplica2012R2EventDetails struct {
	ContainerName       *string `json:"containerName,omitempty"`
	FabricName          *string `json:"fabricName,omitempty"`
	RemoteContainerName *string `json:"remoteContainerName,omitempty"`
	RemoteFabricName    *string `json:"remoteFabricName,omitempty"`

	// Fields inherited from EventProviderSpecificDetails

	InstanceType string `json:"instanceType"`
}

func (s HyperVReplica2012R2EventDetails) EventProviderSpecificDetails() BaseEventProviderSpecificDetailsImpl {
	return BaseEventProviderSpecificDetailsImpl{
		InstanceType: s.InstanceType,
	}
}

var _ json.Marshaler = HyperVReplica2012R2EventDetails{}

func (s HyperVReplica2012R2EventDetails) MarshalJSON() ([]byte, error) {
	type wrapper HyperVReplica2012R2EventDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling HyperVReplica2012R2EventDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling HyperVReplica2012R2EventDetails: %+v", err)
	}

	decoded["instanceType"] = "HyperVReplica2012R2"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling HyperVReplica2012R2EventDetails: %+v", err)
	}

	return encoded, nil
}
