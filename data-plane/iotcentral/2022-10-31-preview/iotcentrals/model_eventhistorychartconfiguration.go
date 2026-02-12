package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TileConfiguration = EventHistoryChartConfiguration{}

type EventHistoryChartConfiguration struct {
	Capabilities *[]TileCapability           `json:"capabilities,omitempty"`
	Devices      []string                    `json:"devices"`
	Format       *TextFormatConfiguration    `json:"format,omitempty"`
	Group        string                      `json:"group"`
	QueryRange   TimeQueryRangeConfiguration `json:"queryRange"`

	// Fields inherited from TileConfiguration

	Type string `json:"type"`
}

func (s EventHistoryChartConfiguration) TileConfiguration() BaseTileConfigurationImpl {
	return BaseTileConfigurationImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = EventHistoryChartConfiguration{}

func (s EventHistoryChartConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper EventHistoryChartConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling EventHistoryChartConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling EventHistoryChartConfiguration: %+v", err)
	}

	decoded["type"] = "eventHistoryChart"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling EventHistoryChartConfiguration: %+v", err)
	}

	return encoded, nil
}
