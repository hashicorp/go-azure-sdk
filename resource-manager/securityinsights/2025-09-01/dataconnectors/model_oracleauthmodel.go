package dataconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CcpAuthConfig = OracleAuthModel{}

type OracleAuthModel struct {
	PemFile           string `json:"pemFile"`
	PublicFingerprint string `json:"publicFingerprint"`
	TenantId          string `json:"tenantId"`
	UserId            string `json:"userId"`

	// Fields inherited from CcpAuthConfig

	Type CcpAuthType `json:"type"`
}

func (s OracleAuthModel) CcpAuthConfig() BaseCcpAuthConfigImpl {
	return BaseCcpAuthConfigImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = OracleAuthModel{}

func (s OracleAuthModel) MarshalJSON() ([]byte, error) {
	type wrapper OracleAuthModel
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling OracleAuthModel: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling OracleAuthModel: %+v", err)
	}

	decoded["type"] = "Oracle"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling OracleAuthModel: %+v", err)
	}

	return encoded, nil
}
