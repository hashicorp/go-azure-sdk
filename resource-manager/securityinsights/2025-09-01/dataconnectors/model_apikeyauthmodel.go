package dataconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CcpAuthConfig = ApiKeyAuthModel{}

type ApiKeyAuthModel struct {
	ApiKey                string  `json:"apiKey"`
	ApiKeyIdentifier      *string `json:"apiKeyIdentifier,omitempty"`
	ApiKeyName            string  `json:"apiKeyName"`
	IsApiKeyInPostPayload *bool   `json:"isApiKeyInPostPayload,omitempty"`

	// Fields inherited from CcpAuthConfig

	Type CcpAuthType `json:"type"`
}

func (s ApiKeyAuthModel) CcpAuthConfig() BaseCcpAuthConfigImpl {
	return BaseCcpAuthConfigImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = ApiKeyAuthModel{}

func (s ApiKeyAuthModel) MarshalJSON() ([]byte, error) {
	type wrapper ApiKeyAuthModel
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ApiKeyAuthModel: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ApiKeyAuthModel: %+v", err)
	}

	decoded["type"] = "APIKey"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ApiKeyAuthModel: %+v", err)
	}

	return encoded, nil
}
