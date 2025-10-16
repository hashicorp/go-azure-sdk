package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ PredictionDriftMetricThresholdBase = CategoricalPredictionDriftMetricThreshold{}

type CategoricalPredictionDriftMetricThreshold struct {
	Metric CategoricalPredictionDriftMetric `json:"metric"`

	// Fields inherited from PredictionDriftMetricThresholdBase

	DataType  MonitoringFeatureDataType `json:"dataType"`
	Threshold *MonitoringThreshold      `json:"threshold,omitempty"`
}

func (s CategoricalPredictionDriftMetricThreshold) PredictionDriftMetricThresholdBase() BasePredictionDriftMetricThresholdBaseImpl {
	return BasePredictionDriftMetricThresholdBaseImpl{
		DataType:  s.DataType,
		Threshold: s.Threshold,
	}
}

var _ json.Marshaler = CategoricalPredictionDriftMetricThreshold{}

func (s CategoricalPredictionDriftMetricThreshold) MarshalJSON() ([]byte, error) {
	type wrapper CategoricalPredictionDriftMetricThreshold
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CategoricalPredictionDriftMetricThreshold: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CategoricalPredictionDriftMetricThreshold: %+v", err)
	}

	decoded["dataType"] = "Categorical"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CategoricalPredictionDriftMetricThreshold: %+v", err)
	}

	return encoded, nil
}
