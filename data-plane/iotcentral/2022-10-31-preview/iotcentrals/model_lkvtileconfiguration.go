package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TileConfiguration = LkvTileConfiguration{}

type LkvTileConfiguration struct {
	Capabilities *[]TileCapability        `json:"capabilities,omitempty"`
	Devices      []string                 `json:"devices"`
	Format       *TextFormatConfiguration `json:"format,omitempty"`
	Group        string                   `json:"group"`
	ShowTrend    *bool                    `json:"showTrend,omitempty"`

	// Fields inherited from TileConfiguration

	Type string `json:"type"`
}

func (s LkvTileConfiguration) TileConfiguration() BaseTileConfigurationImpl {
	return BaseTileConfigurationImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = LkvTileConfiguration{}

func (s LkvTileConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper LkvTileConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling LkvTileConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling LkvTileConfiguration: %+v", err)
	}

	decoded["type"] = "lkv"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling LkvTileConfiguration: %+v", err)
	}

	return encoded, nil
}
