package capabilities

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ComputeModelCapability interface {
	ComputeModelCapability() BaseComputeModelCapabilityImpl
}

var _ ComputeModelCapability = BaseComputeModelCapabilityImpl{}

type BaseComputeModelCapabilityImpl struct {
	Type PoolComputeModelType `json:"type"`
}

func (s BaseComputeModelCapabilityImpl) ComputeModelCapability() BaseComputeModelCapabilityImpl {
	return s
}

var _ ComputeModelCapability = RawComputeModelCapabilityImpl{}

// RawComputeModelCapabilityImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawComputeModelCapabilityImpl struct {
	computeModelCapability BaseComputeModelCapabilityImpl
	Type                   string
	Values                 map[string]interface{}
}

func (s RawComputeModelCapabilityImpl) ComputeModelCapability() BaseComputeModelCapabilityImpl {
	return s.computeModelCapability
}

func (s RawComputeModelCapabilityImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalComputeModelCapabilityImplementation(input []byte) (ComputeModelCapability, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ComputeModelCapability into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Provisioned") {
		var out ProvisionedComputeModelCapability
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ProvisionedComputeModelCapability: %+v", err)
		}
		return out, nil
	}

	var parent BaseComputeModelCapabilityImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseComputeModelCapabilityImpl: %+v", err)
	}

	return RawComputeModelCapabilityImpl{
		computeModelCapability: parent,
		Type:                   value,
		Values:                 temp,
	}, nil

}
