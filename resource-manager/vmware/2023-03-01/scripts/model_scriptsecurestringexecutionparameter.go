package scripts

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ScriptExecutionParameter = ScriptSecureStringExecutionParameter{}

type ScriptSecureStringExecutionParameter struct {
	SecureValue *string `json:"secureValue,omitempty"`

	// Fields inherited from ScriptExecutionParameter

	Name string                       `json:"name"`
	Type ScriptExecutionParameterType `json:"type"`
}

func (s ScriptSecureStringExecutionParameter) ScriptExecutionParameter() BaseScriptExecutionParameterImpl {
	return BaseScriptExecutionParameterImpl{
		Name: s.Name,
		Type: s.Type,
	}
}

var _ json.Marshaler = ScriptSecureStringExecutionParameter{}

func (s ScriptSecureStringExecutionParameter) MarshalJSON() ([]byte, error) {
	type wrapper ScriptSecureStringExecutionParameter
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ScriptSecureStringExecutionParameter: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ScriptSecureStringExecutionParameter: %+v", err)
	}

	decoded["type"] = "SecureValue"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ScriptSecureStringExecutionParameter: %+v", err)
	}

	return encoded, nil
}
