package applicationgroup

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationGroupPolicy interface {
	ApplicationGroupPolicy() BaseApplicationGroupPolicyImpl
}

var _ ApplicationGroupPolicy = BaseApplicationGroupPolicyImpl{}

type BaseApplicationGroupPolicyImpl struct {
	Name string                     `json:"name"`
	Type ApplicationGroupPolicyType `json:"type"`
}

func (s BaseApplicationGroupPolicyImpl) ApplicationGroupPolicy() BaseApplicationGroupPolicyImpl {
	return s
}

var _ ApplicationGroupPolicy = RawApplicationGroupPolicyImpl{}

// RawApplicationGroupPolicyImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawApplicationGroupPolicyImpl struct {
	applicationGroupPolicy BaseApplicationGroupPolicyImpl
	Type                   string
	Values                 map[string]interface{}
}

func (s RawApplicationGroupPolicyImpl) ApplicationGroupPolicy() BaseApplicationGroupPolicyImpl {
	return s.applicationGroupPolicy
}

func UnmarshalApplicationGroupPolicyImplementation(input []byte) (ApplicationGroupPolicy, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ApplicationGroupPolicy into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "ThrottlingPolicy") {
		var out ThrottlingPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ThrottlingPolicy: %+v", err)
		}
		return out, nil
	}

	var parent BaseApplicationGroupPolicyImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseApplicationGroupPolicyImpl: %+v", err)
	}

	return RawApplicationGroupPolicyImpl{
		applicationGroupPolicy: parent,
		Type:                   value,
		Values:                 temp,
	}, nil

}
