package service

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartitionSchemeDescription interface {
}

// RawModeOfTransitImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawPartitionSchemeDescriptionImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalPartitionSchemeDescriptionImplementation(input []byte) (PartitionSchemeDescription, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PartitionSchemeDescription into map[string]interface: %+v", err)
	}

	value, ok := temp["partitionScheme"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Named") {
		var out NamedPartitionSchemeDescription
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NamedPartitionSchemeDescription: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Singleton") {
		var out SingletonPartitionSchemeDescription
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SingletonPartitionSchemeDescription: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "UniformInt64Range") {
		var out UniformInt64RangePartitionSchemeDescription
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UniformInt64RangePartitionSchemeDescription: %+v", err)
		}
		return out, nil
	}

	out := RawPartitionSchemeDescriptionImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
