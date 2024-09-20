package schedule

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SamplingAlgorithm interface {
	SamplingAlgorithm() BaseSamplingAlgorithmImpl
}

var _ SamplingAlgorithm = BaseSamplingAlgorithmImpl{}

type BaseSamplingAlgorithmImpl struct {
	SamplingAlgorithmType SamplingAlgorithmType `json:"samplingAlgorithmType"`
}

func (s BaseSamplingAlgorithmImpl) SamplingAlgorithm() BaseSamplingAlgorithmImpl {
	return s
}

var _ SamplingAlgorithm = RawSamplingAlgorithmImpl{}

// RawSamplingAlgorithmImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSamplingAlgorithmImpl struct {
	samplingAlgorithm BaseSamplingAlgorithmImpl
	Type              string
	Values            map[string]interface{}
}

func (s RawSamplingAlgorithmImpl) SamplingAlgorithm() BaseSamplingAlgorithmImpl {
	return s.samplingAlgorithm
}

func UnmarshalSamplingAlgorithmImplementation(input []byte) (SamplingAlgorithm, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SamplingAlgorithm into map[string]interface: %+v", err)
	}

	value, ok := temp["samplingAlgorithmType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Bayesian") {
		var out BayesianSamplingAlgorithm
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into BayesianSamplingAlgorithm: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Grid") {
		var out GridSamplingAlgorithm
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into GridSamplingAlgorithm: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Random") {
		var out RandomSamplingAlgorithm
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RandomSamplingAlgorithm: %+v", err)
		}
		return out, nil
	}

	var parent BaseSamplingAlgorithmImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSamplingAlgorithmImpl: %+v", err)
	}

	return RawSamplingAlgorithmImpl{
		samplingAlgorithm: parent,
		Type:              value,
		Values:            temp,
	}, nil

}
