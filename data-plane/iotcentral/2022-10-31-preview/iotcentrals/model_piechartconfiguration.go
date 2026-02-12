package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TileConfiguration = PieChartConfiguration{}

type PieChartConfiguration struct {
	Capabilities *[]TileCapability           `json:"capabilities,omitempty"`
	Devices      []string                    `json:"devices"`
	Format       *ChartFormatConfiguration   `json:"format,omitempty"`
	Group        string                      `json:"group"`
	QueryRange   TimeQueryRangeConfiguration `json:"queryRange"`

	// Fields inherited from TileConfiguration

	Type string `json:"type"`
}

func (s PieChartConfiguration) TileConfiguration() BaseTileConfigurationImpl {
	return BaseTileConfigurationImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = PieChartConfiguration{}

func (s PieChartConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper PieChartConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PieChartConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PieChartConfiguration: %+v", err)
	}

	decoded["type"] = "pieChart"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PieChartConfiguration: %+v", err)
	}

	return encoded, nil
}
