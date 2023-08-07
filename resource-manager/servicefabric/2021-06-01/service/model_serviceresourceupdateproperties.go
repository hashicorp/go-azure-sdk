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

// RawModeOfTransitImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawServiceResourceUpdatePropertiesImpl struct {
	Type   string
	Values map[string]interface{}
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

	out := RawServiceResourceUpdatePropertiesImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
