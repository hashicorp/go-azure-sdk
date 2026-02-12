package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TileConfiguration = LineChartConfiguration{}

type LineChartConfiguration struct {
	Capabilities *[]TileCapability         `json:"capabilities,omitempty"`
	Devices      []string                  `json:"devices"`
	Format       *ChartFormatConfiguration `json:"format,omitempty"`
	Group        string                    `json:"group"`
	QueryRange   QueryRangeConfiguration   `json:"queryRange"`

	// Fields inherited from TileConfiguration

	Type string `json:"type"`
}

func (s LineChartConfiguration) TileConfiguration() BaseTileConfigurationImpl {
	return BaseTileConfigurationImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = LineChartConfiguration{}

func (s LineChartConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper LineChartConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling LineChartConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling LineChartConfiguration: %+v", err)
	}

	decoded["type"] = "lineChart"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling LineChartConfiguration: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &LineChartConfiguration{}

func (s *LineChartConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Capabilities *[]TileCapability         `json:"capabilities,omitempty"`
		Devices      []string                  `json:"devices"`
		Format       *ChartFormatConfiguration `json:"format,omitempty"`
		Group        string                    `json:"group"`
		Type         string                    `json:"type"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Capabilities = decoded.Capabilities
	s.Devices = decoded.Devices
	s.Format = decoded.Format
	s.Group = decoded.Group
	s.Type = decoded.Type

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling LineChartConfiguration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["queryRange"]; ok {
		impl, err := UnmarshalQueryRangeConfigurationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'QueryRange' for 'LineChartConfiguration': %+v", err)
		}
		s.QueryRange = impl
	}

	return nil
}
