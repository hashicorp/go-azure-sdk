package experiments

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ChaosExperimentAction interface {
	ChaosExperimentAction() BaseChaosExperimentActionImpl
}

var _ ChaosExperimentAction = BaseChaosExperimentActionImpl{}

type BaseChaosExperimentActionImpl struct {
	Name string               `json:"name"`
	Type ExperimentActionType `json:"type"`
}

func (s BaseChaosExperimentActionImpl) ChaosExperimentAction() BaseChaosExperimentActionImpl {
	return s
}

var _ ChaosExperimentAction = RawChaosExperimentActionImpl{}

// RawChaosExperimentActionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawChaosExperimentActionImpl struct {
	chaosExperimentAction BaseChaosExperimentActionImpl
	Type                  string
	Values                map[string]interface{}
}

func (s RawChaosExperimentActionImpl) ChaosExperimentAction() BaseChaosExperimentActionImpl {
	return s.chaosExperimentAction
}

func UnmarshalChaosExperimentActionImplementation(input []byte) (ChaosExperimentAction, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ChaosExperimentAction into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "continuous") {
		var out ContinuousAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ContinuousAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "delay") {
		var out DelayAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DelayAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "discrete") {
		var out DiscreteAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DiscreteAction: %+v", err)
		}
		return out, nil
	}

	var parent BaseChaosExperimentActionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseChaosExperimentActionImpl: %+v", err)
	}

	return RawChaosExperimentActionImpl{
		chaosExperimentAction: parent,
		Type:                  value,
		Values:                temp,
	}, nil

}
