package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TileConfiguration = MarkdownTileConfiguration{}

type MarkdownTileConfiguration struct {
	Description string  `json:"description"`
	Href        *string `json:"href,omitempty"`
	Image       *string `json:"image,omitempty"`

	// Fields inherited from TileConfiguration

	Type string `json:"type"`
}

func (s MarkdownTileConfiguration) TileConfiguration() BaseTileConfigurationImpl {
	return BaseTileConfigurationImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = MarkdownTileConfiguration{}

func (s MarkdownTileConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper MarkdownTileConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling MarkdownTileConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling MarkdownTileConfiguration: %+v", err)
	}

	decoded["type"] = "markdown"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling MarkdownTileConfiguration: %+v", err)
	}

	return encoded, nil
}
