package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BlobStorageV1DestinationAuth = BlobStorageV1DestinationSystemAssignedManagedIdentityAuth{}

type BlobStorageV1DestinationSystemAssignedManagedIdentityAuth struct {
	ContainerName string `json:"containerName"`
	EndpointUri   string `json:"endpointUri"`

	// Fields inherited from BlobStorageV1DestinationAuth

	Type string `json:"type"`
}

func (s BlobStorageV1DestinationSystemAssignedManagedIdentityAuth) BlobStorageV1DestinationAuth() BaseBlobStorageV1DestinationAuthImpl {
	return BaseBlobStorageV1DestinationAuthImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = BlobStorageV1DestinationSystemAssignedManagedIdentityAuth{}

func (s BlobStorageV1DestinationSystemAssignedManagedIdentityAuth) MarshalJSON() ([]byte, error) {
	type wrapper BlobStorageV1DestinationSystemAssignedManagedIdentityAuth
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BlobStorageV1DestinationSystemAssignedManagedIdentityAuth: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BlobStorageV1DestinationSystemAssignedManagedIdentityAuth: %+v", err)
	}

	decoded["type"] = "systemAssignedManagedIdentity"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BlobStorageV1DestinationSystemAssignedManagedIdentityAuth: %+v", err)
	}

	return encoded, nil
}
