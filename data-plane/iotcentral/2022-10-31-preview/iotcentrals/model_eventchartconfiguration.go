package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TileConfiguration = EventChartConfiguration{}

type EventChartConfiguration struct {
	Capabilities *[]TileCapability           `json:"capabilities,omitempty"`
	Devices      []string                    `json:"devices"`
	Group        string                      `json:"group"`
	QueryRange   TimeQueryRangeConfiguration `json:"queryRange"`

	// Fields inherited from TileConfiguration

	Type string `json:"type"`
}

func (s EventChartConfiguration) TileConfiguration() BaseTileConfigurationImpl {
	return BaseTileConfigurationImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = EventChartConfiguration{}

func (s EventChartConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper EventChartConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EventChartConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EventChartConfiguration: %+v", err)
	}

	decoded["type"] = "eventChart"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EventChartConfiguration: %+v", err)
	}

	return encoded, nil
}
