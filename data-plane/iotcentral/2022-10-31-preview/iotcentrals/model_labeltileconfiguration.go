package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TileConfiguration = LabelTileConfiguration{}

type LabelTileConfiguration struct {
	Text         string            `json:"text"`
	TextSize     *float64          `json:"textSize,omitempty"`
	TextSizeUnit *TileTextSizeUnit `json:"textSizeUnit,omitempty"`
	WordWrap     *bool             `json:"wordWrap,omitempty"`

	// Fields inherited from TileConfiguration

	Type string `json:"type"`
}

func (s LabelTileConfiguration) TileConfiguration() BaseTileConfigurationImpl {
	return BaseTileConfigurationImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = LabelTileConfiguration{}

func (s LabelTileConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper LabelTileConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling LabelTileConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling LabelTileConfiguration: %+v", err)
	}

	decoded["type"] = "label"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling LabelTileConfiguration: %+v", err)
	}

	return encoded, nil
}
