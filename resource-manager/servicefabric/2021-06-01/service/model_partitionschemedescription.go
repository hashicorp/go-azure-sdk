package service

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PartitionSchemeDescription interface {
	PartitionSchemeDescription() BasePartitionSchemeDescriptionImpl
}

var _ PartitionSchemeDescription = BasePartitionSchemeDescriptionImpl{}

type BasePartitionSchemeDescriptionImpl struct {
	PartitionScheme PartitionScheme `json:"partitionScheme"`
}

func (s BasePartitionSchemeDescriptionImpl) PartitionSchemeDescription() BasePartitionSchemeDescriptionImpl {
	return s
}

var _ PartitionSchemeDescription = RawPartitionSchemeDescriptionImpl{}

// RawPartitionSchemeDescriptionImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawPartitionSchemeDescriptionImpl struct {
	partitionSchemeDescription BasePartitionSchemeDescriptionImpl
	Type                       string
	Values                     map[string]interface{}
}

func (s RawPartitionSchemeDescriptionImpl) PartitionSchemeDescription() BasePartitionSchemeDescriptionImpl {
	return s.partitionSchemeDescription
}

func (s RawPartitionSchemeDescriptionImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalPartitionSchemeDescriptionImplementation(input []byte) (PartitionSchemeDescription, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PartitionSchemeDescription into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["partitionScheme"]; ok {
		value = fmt.Sprintf("%v", v)
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

	var parent BasePartitionSchemeDescriptionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePartitionSchemeDescriptionImpl: %+v", err)
	}

	return RawPartitionSchemeDescriptionImpl{
		partitionSchemeDescription: parent,
		Type:                       value,
		Values:                     temp,
	}, nil

}
