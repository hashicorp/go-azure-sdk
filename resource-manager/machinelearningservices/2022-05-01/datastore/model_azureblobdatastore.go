package datastore

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Datastore = AzureBlobDatastore{}

type AzureBlobDatastore struct {
	AccountName                   *string                        `json:"accountName,omitempty"`
	ContainerName                 *string                        `json:"containerName,omitempty"`
	Endpoint                      *string                        `json:"endpoint,omitempty"`
	Protocol                      *string                        `json:"protocol,omitempty"`
	ServiceDataAccessAuthIdentity *ServiceDataAccessAuthIdentity `json:"serviceDataAccessAuthIdentity,omitempty"`

	// Fields inherited from Datastore
	Credentials DatastoreCredentials `json:"credentials"`
	Description *string              `json:"description,omitempty"`
	IsDefault   *bool                `json:"isDefault,omitempty"`
	Properties  *map[string]string   `json:"properties,omitempty"`
	Tags        *map[string]string   `json:"tags,omitempty"`
}

var _ json.Marshaler = AzureBlobDatastore{}

func (s AzureBlobDatastore) MarshalJSON() ([]byte, error) {
	type wrapper AzureBlobDatastore
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureBlobDatastore: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureBlobDatastore: %+v", err)
	}
	decoded["datastoreType"] = "AzureBlob"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureBlobDatastore: %+v", err)
	}

	return encoded, nil
}
