package job

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

// RawEarlyTerminationPolicyImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawEarlyTerminationPolicyImpl struct {
	earlyTerminationPolicy BaseEarlyTerminationPolicyImpl
	Type                   string
	Values                 map[string]interface{}
}

func (s RawEarlyTerminationPolicyImpl) EarlyTerminationPolicy() BaseEarlyTerminationPolicyImpl {
	return s.earlyTerminationPolicy
}

func (s RawEarlyTerminationPolicyImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalEarlyTerminationPolicyImplementation(input []byte) (EarlyTerminationPolicy, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EarlyTerminationPolicy into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["policyType"]; ok {
		value = fmt.Sprintf("%v", v)
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
