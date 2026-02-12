package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TileConfiguration = CommandTileConfiguration{}

type CommandTileConfiguration struct {
	Command string      `json:"command"`
	Device  interface{} `json:"device"`
	Group   string      `json:"group"`

	// Fields inherited from TileConfiguration

	Type string `json:"type"`
}

func (s CommandTileConfiguration) TileConfiguration() BaseTileConfigurationImpl {
	return BaseTileConfigurationImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = CommandTileConfiguration{}

func (s CommandTileConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper CommandTileConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CommandTileConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CommandTileConfiguration: %+v", err)
	}

	decoded["type"] = "command"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CommandTileConfiguration: %+v", err)
	}

	return encoded, nil
}
