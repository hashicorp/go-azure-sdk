package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PredictionDriftMetricThresholdBase = NumericalPredictionDriftMetricThreshold{}

type NumericalPredictionDriftMetricThreshold struct {
	Metric NumericalPredictionDriftMetric `json:"metric"`

	// Fields inherited from PredictionDriftMetricThresholdBase

	DataType  MonitoringFeatureDataType `json:"dataType"`
	Threshold *MonitoringThreshold      `json:"threshold,omitempty"`
}

func (s NumericalPredictionDriftMetricThreshold) PredictionDriftMetricThresholdBase() BasePredictionDriftMetricThresholdBaseImpl {
	return BasePredictionDriftMetricThresholdBaseImpl{
		DataType:  s.DataType,
		Threshold: s.Threshold,
	}
}

var _ json.Marshaler = NumericalPredictionDriftMetricThreshold{}

func (s NumericalPredictionDriftMetricThreshold) MarshalJSON() ([]byte, error) {
	type wrapper NumericalPredictionDriftMetricThreshold
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NumericalPredictionDriftMetricThreshold: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NumericalPredictionDriftMetricThreshold: %+v", err)
	}

	decoded["dataType"] = "Numerical"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NumericalPredictionDriftMetricThreshold: %+v", err)
	}

	return encoded, nil
}
