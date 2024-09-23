package schedule

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TargetLags interface {
	TargetLags() BaseTargetLagsImpl
}

var _ TargetLags = BaseTargetLagsImpl{}

type BaseTargetLagsImpl struct {
	Mode TargetLagsMode `json:"mode"`
}

func (s BaseTargetLagsImpl) TargetLags() BaseTargetLagsImpl {
	return s
}

var _ TargetLags = RawTargetLagsImpl{}

// RawTargetLagsImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawTargetLagsImpl struct {
	targetLags BaseTargetLagsImpl
	Type       string
	Values     map[string]interface{}
}

func (s RawTargetLagsImpl) TargetLags() BaseTargetLagsImpl {
	return s.targetLags
}

func UnmarshalTargetLagsImplementation(input []byte) (TargetLags, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling TargetLags into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["mode"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Auto") {
		var out AutoTargetLags
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AutoTargetLags: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Custom") {
		var out CustomTargetLags
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomTargetLags: %+v", err)
		}
		return out, nil
	}

	var parent BaseTargetLagsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseTargetLagsImpl: %+v", err)
	}

	return RawTargetLagsImpl{
		targetLags: parent,
		Type:       value,
		Values:     temp,
	}, nil

}
