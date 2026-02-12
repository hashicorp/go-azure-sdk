package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Tile struct {
	Configuration TileConfiguration `json:"configuration"`
	DisplayName   string            `json:"displayName"`
	Height        float64           `json:"height"`
	Id            *string           `json:"id,omitempty"`
	Width         float64           `json:"width"`
	X             float64           `json:"x"`
	Y             float64           `json:"y"`
}

var _ json.Unmarshaler = &Tile{}

func (s *Tile) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		DisplayName string  `json:"displayName"`
		Height      float64 `json:"height"`
		Id          *string `json:"id,omitempty"`
		Width       float64 `json:"width"`
		X           float64 `json:"x"`
		Y           float64 `json:"y"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.DisplayName = decoded.DisplayName
	s.Height = decoded.Height
	s.Id = decoded.Id
	s.Width = decoded.Width
	s.X = decoded.X
	s.Y = decoded.Y

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling Tile into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["configuration"]; ok {
		impl, err := UnmarshalTileConfigurationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Configuration' for 'Tile': %+v", err)
		}
		s.Configuration = impl
	}

	return nil
}
