package scripts

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ScriptExecutionParameter = PSCredentialExecutionParameter{}

type PSCredentialExecutionParameter struct {
	Password *string `json:"password,omitempty"`
	Username *string `json:"username,omitempty"`

	// Fields inherited from ScriptExecutionParameter
	Name string `json:"name"`
}

var _ json.Marshaler = PSCredentialExecutionParameter{}

func (s PSCredentialExecutionParameter) MarshalJSON() ([]byte, error) {
	type wrapper PSCredentialExecutionParameter
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PSCredentialExecutionParameter: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PSCredentialExecutionParameter: %+v", err)
	}
	decoded["type"] = "Credential"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PSCredentialExecutionParameter: %+v", err)
	}

	return encoded, nil
}
