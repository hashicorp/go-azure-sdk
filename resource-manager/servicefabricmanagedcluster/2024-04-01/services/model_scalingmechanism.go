package services

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScalingMechanism interface {
	ScalingMechanism() BaseScalingMechanismImpl
}

var _ ScalingMechanism = BaseScalingMechanismImpl{}

type BaseScalingMechanismImpl struct {
	Kind ServiceScalingMechanismKind `json:"kind"`
}

func (s BaseScalingMechanismImpl) ScalingMechanism() BaseScalingMechanismImpl {
	return s
}

var _ ScalingMechanism = RawScalingMechanismImpl{}

// RawScalingMechanismImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawScalingMechanismImpl struct {
	scalingMechanism BaseScalingMechanismImpl
	Type             string
	Values           map[string]interface{}
}

func (s RawScalingMechanismImpl) ScalingMechanism() BaseScalingMechanismImpl {
	return s.scalingMechanism
}

func (s RawScalingMechanismImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalScalingMechanismImplementation(input []byte) (ScalingMechanism, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ScalingMechanism into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["kind"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "AddRemoveIncrementalNamedPartition") {
		var out AddRemoveIncrementalNamedPartitionScalingMechanism
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AddRemoveIncrementalNamedPartitionScalingMechanism: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ScalePartitionInstanceCount") {
		var out PartitionInstanceCountScaleMechanism
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PartitionInstanceCountScaleMechanism: %+v", err)
		}
		return out, nil
	}

	var parent BaseScalingMechanismImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseScalingMechanismImpl: %+v", err)
	}

	return RawScalingMechanismImpl{
		scalingMechanism: parent,
		Type:             value,
		Values:           temp,
	}, nil

}
