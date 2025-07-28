package dataconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CcpAuthConfig = SessionAuthModel{}

type SessionAuthModel struct {
	Headers                 *map[string]string      `json:"headers,omitempty"`
	IsPostPayloadJson       *bool                   `json:"isPostPayloadJson,omitempty"`
	Password                map[string]string       `json:"password"`
	QueryParameters         *map[string]interface{} `json:"queryParameters,omitempty"`
	SessionIdName           *string                 `json:"sessionIdName,omitempty"`
	SessionLoginRequestUri  *string                 `json:"sessionLoginRequestUri,omitempty"`
	SessionTimeoutInMinutes *int64                  `json:"sessionTimeoutInMinutes,omitempty"`
	UserName                map[string]string       `json:"userName"`

	// Fields inherited from CcpAuthConfig

	Type CcpAuthType `json:"type"`
}

func (s SessionAuthModel) CcpAuthConfig() BaseCcpAuthConfigImpl {
	return BaseCcpAuthConfigImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = SessionAuthModel{}

func (s SessionAuthModel) MarshalJSON() ([]byte, error) {
	type wrapper SessionAuthModel
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SessionAuthModel: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SessionAuthModel: %+v", err)
	}

	decoded["type"] = "Session"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SessionAuthModel: %+v", err)
	}

	return encoded, nil
}
