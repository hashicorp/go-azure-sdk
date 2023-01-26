package service

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServiceResourceUpdateProperties interface {
}

func unmarshalServiceResourceUpdatePropertiesImplementation(input []byte) (ServiceResourceUpdateProperties, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ServiceResourceUpdateProperties into map[string]interface: %+v", err)
	}

	value, ok := temp["serviceKind"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Stateful") {
		var out StatefulServiceUpdateProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StatefulServiceUpdateProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Stateless") {
		var out StatelessServiceUpdateProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StatelessServiceUpdateProperties: %+v", err)
		}
		return out, nil
	}

	type RawServiceResourceUpdatePropertiesImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawServiceResourceUpdatePropertiesImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
