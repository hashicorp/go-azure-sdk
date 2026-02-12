package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TileConfiguration = MapPropertyConfiguration{}

type MapPropertyConfiguration struct {
	Capabilities *[]TileCapability `json:"capabilities,omitempty"`
	Devices      []string          `json:"devices"`
	Group        string            `json:"group"`
	ZoomLevel    *float64          `json:"zoomLevel,omitempty"`

	// Fields inherited from TileConfiguration

	Type string `json:"type"`
}

func (s MapPropertyConfiguration) TileConfiguration() BaseTileConfigurationImpl {
	return BaseTileConfigurationImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = MapPropertyConfiguration{}

func (s MapPropertyConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper MapPropertyConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MapPropertyConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MapPropertyConfiguration: %+v", err)
	}

	decoded["type"] = "mapProperty"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MapPropertyConfiguration: %+v", err)
	}

	return encoded, nil
}
