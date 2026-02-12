package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TileConfiguration = ImageTileConfiguration{}

type ImageTileConfiguration struct {
	Format *ImageTileConfigurationFormat `json:"format,omitempty"`
	Href   *string                       `json:"href,omitempty"`
	Image  *string                       `json:"image,omitempty"`

	// Fields inherited from TileConfiguration

	Type string `json:"type"`
}

func (s ImageTileConfiguration) TileConfiguration() BaseTileConfigurationImpl {
	return BaseTileConfigurationImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = ImageTileConfiguration{}

func (s ImageTileConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper ImageTileConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ImageTileConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ImageTileConfiguration: %+v", err)
	}

	decoded["type"] = "image"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ImageTileConfiguration: %+v", err)
	}

	return encoded, nil
}
