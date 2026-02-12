package pipelines

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Activity interface {
	Activity() BaseActivityImpl
}

var _ Activity = BaseActivityImpl{}

type BaseActivityImpl struct {
	DependsOn      *[]ActivityDependency `json:"dependsOn,omitempty"`
	Description    *string               `json:"description,omitempty"`
	Name           string                `json:"name"`
	Type           string                `json:"type"`
	UserProperties *[]UserProperty       `json:"userProperties,omitempty"`
}

func (s BaseActivityImpl) Activity() BaseActivityImpl {
	return s
}

var _ Activity = RawActivityImpl{}

// RawActivityImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawActivityImpl struct {
	activity BaseActivityImpl
	Type     string
	Values   map[string]interface{}
}

func (s RawActivityImpl) Activity() BaseActivityImpl {
	return s.activity
}

func UnmarshalActivityImplementation(input []byte) (Activity, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Activity into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Container") {
		var out ControlActivity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ControlActivity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Execution") {
		var out ExecutionActivity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExecutionActivity: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SqlPoolStoredProcedure") {
		var out SqlPoolStoredProcedureActivity
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SqlPoolStoredProcedureActivity: %+v", err)
		}
		return out, nil
	}

	var parent BaseActivityImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseActivityImpl: %+v", err)
	}

	return RawActivityImpl{
		activity: parent,
		Type:     value,
		Values:   temp,
	}, nil

}
