package services

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScalingTrigger interface {
	ScalingTrigger() BaseScalingTriggerImpl
}

var _ ScalingTrigger = BaseScalingTriggerImpl{}

type BaseScalingTriggerImpl struct {
	Kind ServiceScalingTriggerKind `json:"kind"`
}

func (s BaseScalingTriggerImpl) ScalingTrigger() BaseScalingTriggerImpl {
	return s
}

var _ ScalingTrigger = RawScalingTriggerImpl{}

// RawScalingTriggerImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawScalingTriggerImpl struct {
	scalingTrigger BaseScalingTriggerImpl
	Type           string
	Values         map[string]interface{}
}

func (s RawScalingTriggerImpl) ScalingTrigger() BaseScalingTriggerImpl {
	return s.scalingTrigger
}

func (s RawScalingTriggerImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalScalingTriggerImplementation(input []byte) (ScalingTrigger, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ScalingTrigger into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["kind"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "AveragePartitionLoadTrigger") {
		var out AveragePartitionLoadScalingTrigger
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AveragePartitionLoadScalingTrigger: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AverageServiceLoadTrigger") {
		var out AverageServiceLoadScalingTrigger
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AverageServiceLoadScalingTrigger: %+v", err)
		}
		return out, nil
	}

	var parent BaseScalingTriggerImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseScalingTriggerImpl: %+v", err)
	}

	return RawScalingTriggerImpl{
		scalingTrigger: parent,
		Type:           value,
		Values:         temp,
	}, nil

}
