package datastore

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ Datastore = AzureFileDatastore{}

type AzureFileDatastore struct {
	AccountName                   string                         `json:"accountName"`
	Endpoint                      *string                        `json:"endpoint,omitempty"`
	FileShareName                 string                         `json:"fileShareName"`
	Protocol                      *string                        `json:"protocol,omitempty"`
	ServiceDataAccessAuthIdentity *ServiceDataAccessAuthIdentity `json:"serviceDataAccessAuthIdentity,omitempty"`

	// Fields inherited from Datastore
	Credentials DatastoreCredentials `json:"credentials"`
	Description *string              `json:"description,omitempty"`
	IsDefault   *bool                `json:"isDefault,omitempty"`
	Properties  *map[string]string   `json:"properties,omitempty"`
	Tags        *map[string]string   `json:"tags,omitempty"`
}

var _ json.Marshaler = AzureFileDatastore{}

func (s AzureFileDatastore) MarshalJSON() ([]byte, error) {
	type wrapper AzureFileDatastore
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureFileDatastore: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureFileDatastore: %+v", err)
	}
	decoded["datastoreType"] = "AzureFile"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureFileDatastore: %+v", err)
	}

	return encoded, nil
}
