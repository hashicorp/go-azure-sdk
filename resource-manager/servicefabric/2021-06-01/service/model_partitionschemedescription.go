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

	type RawPartitionSchemeDescriptionImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawPartitionSchemeDescriptionImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
