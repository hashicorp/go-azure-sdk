package triggers

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DependencyReference interface {
	DependencyReference() BaseDependencyReferenceImpl
}

var _ DependencyReference = BaseDependencyReferenceImpl{}

type BaseDependencyReferenceImpl struct {
	Type string `json:"type"`
}

func (s BaseDependencyReferenceImpl) DependencyReference() BaseDependencyReferenceImpl {
	return s
}

var _ DependencyReference = RawDependencyReferenceImpl{}

// RawDependencyReferenceImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDependencyReferenceImpl struct {
	dependencyReference BaseDependencyReferenceImpl
	Type                string
	Values              map[string]interface{}
}

func (s RawDependencyReferenceImpl) DependencyReference() BaseDependencyReferenceImpl {
	return s.dependencyReference
}

func UnmarshalDependencyReferenceImplementation(input []byte) (DependencyReference, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DependencyReference into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "SelfDependencyTumblingWindowTriggerReference") {
		var out SelfDependencyTumblingWindowTriggerReference
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SelfDependencyTumblingWindowTriggerReference: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "TriggerDependencyReference") {
		var out TriggerDependencyReference
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TriggerDependencyReference: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "TumblingWindowTriggerDependencyReference") {
		var out TumblingWindowTriggerDependencyReference
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TumblingWindowTriggerDependencyReference: %+v", err)
		}
		return out, nil
	}

	var parent BaseDependencyReferenceImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDependencyReferenceImpl: %+v", err)
	}

	return RawDependencyReferenceImpl{
		dependencyReference: parent,
		Type:                value,
		Values:              temp,
	}, nil

}
