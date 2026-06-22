package service

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Partition interface {
	Partition() BasePartitionImpl
}

var _ Partition = BasePartitionImpl{}

type BasePartitionImpl struct {
	PartitionScheme PartitionScheme `json:"partitionScheme"`
}

func (s BasePartitionImpl) Partition() BasePartitionImpl {
	return s
}

var _ Partition = RawPartitionImpl{}

// RawPartitionImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawPartitionImpl struct {
	partition BasePartitionImpl
	Type      string
	Values    map[string]interface{}
}

func (s RawPartitionImpl) Partition() BasePartitionImpl {
	return s.partition
}

func (s RawPartitionImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalPartitionImplementation(input []byte) (Partition, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Partition into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["partitionScheme"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Named") {
		var out NamedPartitionScheme
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NamedPartitionScheme: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Singleton") {
		var out SingletonPartitionScheme
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SingletonPartitionScheme: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "UniformInt64Range") {
		var out UniformInt64RangePartitionScheme
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into UniformInt64RangePartitionScheme: %+v", err)
		}
		return out, nil
	}

	var parent BasePartitionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePartitionImpl: %+v", err)
	}

	return RawPartitionImpl{
		partition: parent,
		Type:      value,
		Values:    temp,
	}, nil

}
