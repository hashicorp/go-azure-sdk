package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TileConfiguration = DataExplorerTileConfiguration{}

type DataExplorerTileConfiguration struct {
	Query      *string                     `json:"query,omitempty"`
	QueryRange TimeQueryRangeConfiguration `json:"queryRange"`

	// Fields inherited from TileConfiguration

	Type string `json:"type"`
}

func (s DataExplorerTileConfiguration) TileConfiguration() BaseTileConfigurationImpl {
	return BaseTileConfigurationImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = DataExplorerTileConfiguration{}

func (s DataExplorerTileConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper DataExplorerTileConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DataExplorerTileConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DataExplorerTileConfiguration: %+v", err)
	}

	decoded["type"] = "dataExplorer"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DataExplorerTileConfiguration: %+v", err)
	}

	return encoded, nil
}
