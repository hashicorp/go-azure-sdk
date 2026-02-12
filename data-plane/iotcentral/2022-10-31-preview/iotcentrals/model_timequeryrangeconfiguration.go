package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ QueryRangeConfiguration = TimeQueryRangeConfiguration{}

type TimeQueryRangeConfiguration struct {
	Duration   TileTimeRangeDuration    `json:"duration"`
	Resolution *TileTimeRangeResolution `json:"resolution,omitempty"`

	// Fields inherited from QueryRangeConfiguration

	Type string `json:"type"`
}

func (s TimeQueryRangeConfiguration) QueryRangeConfiguration() BaseQueryRangeConfigurationImpl {
	return BaseQueryRangeConfigurationImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = TimeQueryRangeConfiguration{}

func (s TimeQueryRangeConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper TimeQueryRangeConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling TimeQueryRangeConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling TimeQueryRangeConfiguration: %+v", err)
	}

	decoded["type"] = "time"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling TimeQueryRangeConfiguration: %+v", err)
	}

	return encoded, nil
}
