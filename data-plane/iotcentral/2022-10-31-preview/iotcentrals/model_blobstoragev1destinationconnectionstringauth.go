package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BlobStorageV1DestinationAuth = BlobStorageV1DestinationConnectionStringAuth{}

type BlobStorageV1DestinationConnectionStringAuth struct {
	ConnectionString string `json:"connectionString"`
	ContainerName    string `json:"containerName"`

	// Fields inherited from BlobStorageV1DestinationAuth

	Type string `json:"type"`
}

func (s BlobStorageV1DestinationConnectionStringAuth) BlobStorageV1DestinationAuth() BaseBlobStorageV1DestinationAuthImpl {
	return BaseBlobStorageV1DestinationAuthImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = BlobStorageV1DestinationConnectionStringAuth{}

func (s BlobStorageV1DestinationConnectionStringAuth) MarshalJSON() ([]byte, error) {
	type wrapper BlobStorageV1DestinationConnectionStringAuth
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BlobStorageV1DestinationConnectionStringAuth: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BlobStorageV1DestinationConnectionStringAuth: %+v", err)
	}

	decoded["type"] = "connectionString"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BlobStorageV1DestinationConnectionStringAuth: %+v", err)
	}

	return encoded, nil
}
