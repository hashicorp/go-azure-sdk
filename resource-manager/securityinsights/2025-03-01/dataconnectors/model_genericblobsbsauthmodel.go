package dataconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CcpAuthConfig = GenericBlobSbsAuthModel{}

type GenericBlobSbsAuthModel struct {
	CredentialsConfig               *map[string]string `json:"credentialsConfig,omitempty"`
	StorageAccountCredentialsConfig *map[string]string `json:"storageAccountCredentialsConfig,omitempty"`

	// Fields inherited from CcpAuthConfig

	Type CcpAuthType `json:"type"`
}

func (s GenericBlobSbsAuthModel) CcpAuthConfig() BaseCcpAuthConfigImpl {
	return BaseCcpAuthConfigImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = GenericBlobSbsAuthModel{}

func (s GenericBlobSbsAuthModel) MarshalJSON() ([]byte, error) {
	type wrapper GenericBlobSbsAuthModel
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling GenericBlobSbsAuthModel: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling GenericBlobSbsAuthModel: %+v", err)
	}

	decoded["type"] = "ServiceBus"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling GenericBlobSbsAuthModel: %+v", err)
	}

	return encoded, nil
}
