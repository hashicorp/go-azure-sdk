package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TileConfiguration = ExternalContentTileConfiguration{}

type ExternalContentTileConfiguration struct {
	SourceURL string `json:"sourceUrl"`

	// Fields inherited from TileConfiguration

	Type string `json:"type"`
}

func (s ExternalContentTileConfiguration) TileConfiguration() BaseTileConfigurationImpl {
	return BaseTileConfigurationImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = ExternalContentTileConfiguration{}

func (s ExternalContentTileConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper ExternalContentTileConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExternalContentTileConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExternalContentTileConfiguration: %+v", err)
	}

	decoded["type"] = "externalContent"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExternalContentTileConfiguration: %+v", err)
	}

	return encoded, nil
}
