package scripts

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ScriptExecutionParameter = ScriptStringExecutionParameter{}

type ScriptStringExecutionParameter struct {
	Value *string `json:"value,omitempty"`

	// Fields inherited from ScriptExecutionParameter
	Name string `json:"name"`
}

var _ json.Marshaler = ScriptStringExecutionParameter{}

func (s ScriptStringExecutionParameter) MarshalJSON() ([]byte, error) {
	type wrapper ScriptStringExecutionParameter
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ScriptStringExecutionParameter: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ScriptStringExecutionParameter: %+v", err)
	}
	decoded["type"] = "Value"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ScriptStringExecutionParameter: %+v", err)
	}

	return encoded, nil
}
