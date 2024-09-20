package schedule

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PredictionDriftMetricThresholdBase interface {
	PredictionDriftMetricThresholdBase() BasePredictionDriftMetricThresholdBaseImpl
}

var _ PredictionDriftMetricThresholdBase = BasePredictionDriftMetricThresholdBaseImpl{}

type BasePredictionDriftMetricThresholdBaseImpl struct {
	DataType  MonitoringFeatureDataType `json:"dataType"`
	Threshold *MonitoringThreshold      `json:"threshold,omitempty"`
}

func (s BasePredictionDriftMetricThresholdBaseImpl) PredictionDriftMetricThresholdBase() BasePredictionDriftMetricThresholdBaseImpl {
	return s
}

var _ PredictionDriftMetricThresholdBase = RawPredictionDriftMetricThresholdBaseImpl{}

// RawPredictionDriftMetricThresholdBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawPredictionDriftMetricThresholdBaseImpl struct {
	predictionDriftMetricThresholdBase BasePredictionDriftMetricThresholdBaseImpl
	Type                               string
	Values                             map[string]interface{}
}

func (s RawPredictionDriftMetricThresholdBaseImpl) PredictionDriftMetricThresholdBase() BasePredictionDriftMetricThresholdBaseImpl {
	return s.predictionDriftMetricThresholdBase
}

func UnmarshalPredictionDriftMetricThresholdBaseImplementation(input []byte) (PredictionDriftMetricThresholdBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling PredictionDriftMetricThresholdBase into map[string]interface: %+v", err)
	}

	value, ok := temp["dataType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Categorical") {
		var out CategoricalPredictionDriftMetricThreshold
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CategoricalPredictionDriftMetricThreshold: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Numerical") {
		var out NumericalPredictionDriftMetricThreshold
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NumericalPredictionDriftMetricThreshold: %+v", err)
		}
		return out, nil
	}

	var parent BasePredictionDriftMetricThresholdBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BasePredictionDriftMetricThresholdBaseImpl: %+v", err)
	}

	return RawPredictionDriftMetricThresholdBaseImpl{
		predictionDriftMetricThresholdBase: parent,
		Type:                               value,
		Values:                             temp,
	}, nil

}
