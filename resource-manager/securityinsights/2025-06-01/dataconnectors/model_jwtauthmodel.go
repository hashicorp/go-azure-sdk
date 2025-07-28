package dataconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CcpAuthConfig = JwtAuthModel{}

type JwtAuthModel struct {
	Headers                 *map[string]string `json:"headers,omitempty"`
	IsCredentialsInHeaders  *bool              `json:"isCredentialsInHeaders,omitempty"`
	IsJsonRequest           *bool              `json:"isJsonRequest,omitempty"`
	Password                map[string]string  `json:"password"`
	QueryParameters         *map[string]string `json:"queryParameters,omitempty"`
	RequestTimeoutInSeconds *int64             `json:"requestTimeoutInSeconds,omitempty"`
	TokenEndpoint           string             `json:"tokenEndpoint"`
	UserName                map[string]string  `json:"userName"`

	// Fields inherited from CcpAuthConfig

	Type CcpAuthType `json:"type"`
}

func (s JwtAuthModel) CcpAuthConfig() BaseCcpAuthConfigImpl {
	return BaseCcpAuthConfigImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = JwtAuthModel{}

func (s JwtAuthModel) MarshalJSON() ([]byte, error) {
	type wrapper JwtAuthModel
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling JwtAuthModel: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling JwtAuthModel: %+v", err)
	}

	decoded["type"] = "JwtToken"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling JwtAuthModel: %+v", err)
	}

	return encoded, nil
}
