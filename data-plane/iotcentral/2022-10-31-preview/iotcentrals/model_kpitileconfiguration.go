package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TileConfiguration = KpiTileConfiguration{}

type KpiTileConfiguration struct {
	Capabilities *[]TileCapability           `json:"capabilities,omitempty"`
	Devices      []string                    `json:"devices"`
	Format       *TextFormatConfiguration    `json:"format,omitempty"`
	Group        string                      `json:"group"`
	QueryRange   TimeQueryRangeConfiguration `json:"queryRange"`

	// Fields inherited from TileConfiguration

	Type string `json:"type"`
}

func (s KpiTileConfiguration) TileConfiguration() BaseTileConfigurationImpl {
	return BaseTileConfigurationImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = KpiTileConfiguration{}

func (s KpiTileConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper KpiTileConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling KpiTileConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling KpiTileConfiguration: %+v", err)
	}

	decoded["type"] = "kpi"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling KpiTileConfiguration: %+v", err)
	}

	return encoded, nil
}
