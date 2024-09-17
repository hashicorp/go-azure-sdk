package schedule

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EarlyTerminationPolicy interface {
	EarlyTerminationPolicy() BaseEarlyTerminationPolicyImpl
}

var _ EarlyTerminationPolicy = BaseEarlyTerminationPolicyImpl{}

type BaseEarlyTerminationPolicyImpl struct {
	DelayEvaluation    *int64                     `json:"delayEvaluation,omitempty"`
	EvaluationInterval *int64                     `json:"evaluationInterval,omitempty"`
	PolicyType         EarlyTerminationPolicyType `json:"policyType"`
}

func (s BaseEarlyTerminationPolicyImpl) EarlyTerminationPolicy() BaseEarlyTerminationPolicyImpl {
	return s
}

var _ EarlyTerminationPolicy = RawEarlyTerminationPolicyImpl{}

// RawEarlyTerminationPolicyImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawEarlyTerminationPolicyImpl struct {
	earlyTerminationPolicy BaseEarlyTerminationPolicyImpl
	Type                   string
	Values                 map[string]interface{}
}

func (s RawEarlyTerminationPolicyImpl) EarlyTerminationPolicy() BaseEarlyTerminationPolicyImpl {
	return s.earlyTerminationPolicy
}

func UnmarshalEarlyTerminationPolicyImplementation(input []byte) (EarlyTerminationPolicy, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EarlyTerminationPolicy into map[string]interface: %+v", err)
	}

	value, ok := temp["policyType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Bandit") {
		var out BanditPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BanditPolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "MedianStopping") {
		var out MedianStoppingPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MedianStoppingPolicy: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "TruncationSelection") {
		var out TruncationSelectionPolicy
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TruncationSelectionPolicy: %+v", err)
		}
		return out, nil
	}

	var parent BaseEarlyTerminationPolicyImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEarlyTerminationPolicyImpl: %+v", err)
	}

	return RawEarlyTerminationPolicyImpl{
		earlyTerminationPolicy: parent,
		Type:                   value,
		Values:                 temp,
	}, nil

}
