package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TileConfiguration = PropertyTileConfiguration{}

type PropertyTileConfiguration struct {
	Capabilities *[]TileCapability        `json:"capabilities,omitempty"`
	Devices      []string                 `json:"devices"`
	Format       *TextFormatConfiguration `json:"format,omitempty"`
	Group        string                   `json:"group"`

	// Fields inherited from TileConfiguration

	Type string `json:"type"`
}

func (s PropertyTileConfiguration) TileConfiguration() BaseTileConfigurationImpl {
	return BaseTileConfigurationImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = PropertyTileConfiguration{}

func (s PropertyTileConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper PropertyTileConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PropertyTileConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PropertyTileConfiguration: %+v", err)
	}

	decoded["type"] = "property"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PropertyTileConfiguration: %+v", err)
	}

	return encoded, nil
}
