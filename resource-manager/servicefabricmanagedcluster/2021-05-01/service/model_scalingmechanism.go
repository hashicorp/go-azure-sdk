package service

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

// RawScalingMechanismImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawScalingMechanismImpl struct {
	scalingMechanism BaseScalingMechanismImpl
	Type             string
	Values           map[string]interface{}
}

func (s RawScalingMechanismImpl) ScalingMechanism() BaseScalingMechanismImpl {
	return s.scalingMechanism
}

func UnmarshalScalingMechanismImplementation(input []byte) (ScalingMechanism, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ScalingMechanism into map[string]interface: %+v", err)
	}

	value, ok := temp["kind"].(string)
	if !ok {
		return nil, nil
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
