package iotcentrals

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ QueryRangeConfiguration = CountQueryRangeConfiguration{}

type CountQueryRangeConfiguration struct {
	Count int64 `json:"count"`

	// Fields inherited from QueryRangeConfiguration

	Type string `json:"type"`
}

func (s CountQueryRangeConfiguration) QueryRangeConfiguration() BaseQueryRangeConfigurationImpl {
	return BaseQueryRangeConfigurationImpl{
		Type: s.Type,
	}
}

var _ json.Marshaler = CountQueryRangeConfiguration{}

func (s CountQueryRangeConfiguration) MarshalJSON() ([]byte, error) {
	type wrapper CountQueryRangeConfiguration
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CountQueryRangeConfiguration: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CountQueryRangeConfiguration: %+v", err)
	}

	decoded["type"] = "count"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CountQueryRangeConfiguration: %+v", err)
	}

	return encoded, nil
}
