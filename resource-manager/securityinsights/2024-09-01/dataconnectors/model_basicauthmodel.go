package dataconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CcpAuthConfig = BasicAuthModel{}

type BasicAuthModel struct {
	Password string `json:"password"`
	UserName string `json:"userName"`

	// Fields inherited from CcpAuthConfig

	Type CcpAuthType `json:"type"`
}

func (s BasicAuthModel) CcpAuthConfig() BaseCcpAuthConfigImpl {
	return BaseCcpAuthConfigImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = BasicAuthModel{}

func (s BasicAuthModel) MarshalJSON() ([]byte, error) {
	type wrapper BasicAuthModel
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling BasicAuthModel: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling BasicAuthModel: %+v", err)
	}

	decoded["type"] = "Basic"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling BasicAuthModel: %+v", err)
	}

	return encoded, nil
}
