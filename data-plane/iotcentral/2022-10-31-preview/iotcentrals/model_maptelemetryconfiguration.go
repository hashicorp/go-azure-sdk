package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TileConfiguration = MapTelemetryConfiguration{}

type MapTelemetryConfiguration struct {
	Capabilities *[]TileCapability `json:"capabilities,omitempty"`
	Devices      []string          `json:"devices"`
	Group        string            `json:"group"`
	ZoomLevel    *float64          `json:"zoomLevel,omitempty"`

	// Fields inherited from TileConfiguration

	Type string `json:"type"`
}

func (s MapTelemetryConfiguration) TileConfiguration() BaseTileConfigurationImpl {
	return BaseTileConfigurationImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = MapTelemetryConfiguration{}

func (s MapTelemetryConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper MapTelemetryConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MapTelemetryConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MapTelemetryConfiguration: %+v", err)
	}

	decoded["type"] = "mapTelemetry"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MapTelemetryConfiguration: %+v", err)
	}

	return encoded, nil
}
