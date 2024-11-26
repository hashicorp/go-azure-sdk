package dataconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CcpAuthConfig = AWSAuthModel{}

type AWSAuthModel struct {
	ExternalId *string `json:"externalId,omitempty"`
	RoleArn    string  `json:"roleArn"`

	// Fields inherited from CcpAuthConfig

	Type CcpAuthType `json:"type"`
}

func (s AWSAuthModel) CcpAuthConfig() BaseCcpAuthConfigImpl {
	return BaseCcpAuthConfigImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = AWSAuthModel{}

func (s AWSAuthModel) MarshalJSON() ([]byte, error) {
	type wrapper AWSAuthModel
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AWSAuthModel: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AWSAuthModel: %+v", err)
	}

	decoded["type"] = "AWS"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AWSAuthModel: %+v", err)
	}

	return encoded, nil
}
