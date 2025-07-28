package dataconnectors

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ CcpAuthConfig = NoneAuthModel{}

type NoneAuthModel struct {

	// Fields inherited from CcpAuthConfig

	Type CcpAuthType `json:"type"`
}

func (s NoneAuthModel) CcpAuthConfig() BaseCcpAuthConfigImpl {
	return BaseCcpAuthConfigImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = NoneAuthModel{}

func (s NoneAuthModel) MarshalJSON() ([]byte, error) {
	type wrapper NoneAuthModel
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NoneAuthModel: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NoneAuthModel: %+v", err)
	}

	decoded["type"] = "None"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NoneAuthModel: %+v", err)
	}

	return encoded, nil
}
