package modelversion

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BaseEnvironmentSource interface {
}

func unmarshalBaseEnvironmentSourceImplementation(input []byte) (BaseEnvironmentSource, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling BaseEnvironmentSource into map[string]interface: %+v", err)
	}

	value, ok := temp["baseEnvironmentSourceType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "EnvironmentAsset") {
		var out BaseEnvironmentId
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BaseEnvironmentId: %+v", err)
		}
		return out, nil
	}

	type RawBaseEnvironmentSourceImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawBaseEnvironmentSourceImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
