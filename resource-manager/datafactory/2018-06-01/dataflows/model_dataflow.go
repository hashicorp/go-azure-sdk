package dataflows

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataFlow interface {
}

// RawDataFlowImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDataFlowImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalDataFlowImplementation(input []byte) (DataFlow, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DataFlow into map[string]interface: %+v", err)
	}

	value, ok := temp["type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Flowlet") {
		var out Flowlet
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into Flowlet: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "MappingDataFlow") {
		var out MappingDataFlow
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MappingDataFlow: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "WranglingDataFlow") {
		var out WranglingDataFlow
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into WranglingDataFlow: %+v", err)
		}
		return out, nil
	}

	out := RawDataFlowImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
