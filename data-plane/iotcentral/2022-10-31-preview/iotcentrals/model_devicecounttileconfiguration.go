package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TileConfiguration = DeviceCountTileConfiguration{}

type DeviceCountTileConfiguration struct {
	Format *TextFormatConfiguration `json:"format,omitempty"`
	Group  *string                  `json:"group,omitempty"`

	// Fields inherited from TileConfiguration

	Type string `json:"type"`
}

func (s DeviceCountTileConfiguration) TileConfiguration() BaseTileConfigurationImpl {
	return BaseTileConfigurationImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = DeviceCountTileConfiguration{}

func (s DeviceCountTileConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper DeviceCountTileConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DeviceCountTileConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DeviceCountTileConfiguration: %+v", err)
	}

	decoded["type"] = "deviceCount"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DeviceCountTileConfiguration: %+v", err)
	}

	return encoded, nil
}
