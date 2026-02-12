package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TileConfiguration = HeatMapConfiguration{}

type HeatMapConfiguration struct {
	Capabilities *[]TileCapability           `json:"capabilities,omitempty"`
	Devices      []string                    `json:"devices"`
	Format       *ChartFormatConfiguration   `json:"format,omitempty"`
	Group        string                      `json:"group"`
	QueryRange   TimeQueryRangeConfiguration `json:"queryRange"`

	// Fields inherited from TileConfiguration

	Type string `json:"type"`
}

func (s HeatMapConfiguration) TileConfiguration() BaseTileConfigurationImpl {
	return BaseTileConfigurationImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = HeatMapConfiguration{}

func (s HeatMapConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper HeatMapConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling HeatMapConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling HeatMapConfiguration: %+v", err)
	}

	decoded["type"] = "heatMapChart"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling HeatMapConfiguration: %+v", err)
	}

	return encoded, nil
}
