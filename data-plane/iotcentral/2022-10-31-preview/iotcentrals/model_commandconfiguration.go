package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TileConfiguration = CommandConfiguration{}

type CommandConfiguration struct {
	Command string `json:"command"`
	Device  string `json:"device"`

	// Fields inherited from TileConfiguration

	Type string `json:"type"`
}

func (s CommandConfiguration) TileConfiguration() BaseTileConfigurationImpl {
	return BaseTileConfigurationImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = CommandConfiguration{}

func (s CommandConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper CommandConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CommandConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CommandConfiguration: %+v", err)
	}

	decoded["type"] = "command"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CommandConfiguration: %+v", err)
	}

	return encoded, nil
}
