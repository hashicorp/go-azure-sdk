package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataDriftMetricThresholdBase = CategoricalDataDriftMetricThreshold{}

type CategoricalDataDriftMetricThreshold struct {
	Metric CategoricalDataDriftMetric `json:"metric"`

	// Fields inherited from DataDriftMetricThresholdBase

	DataType  MonitoringFeatureDataType `json:"dataType"`
	Threshold *MonitoringThreshold      `json:"threshold,omitempty"`
}

func (s CategoricalDataDriftMetricThreshold) DataDriftMetricThresholdBase() BaseDataDriftMetricThresholdBaseImpl {
	return BaseDataDriftMetricThresholdBaseImpl{
		DataType:  s.DataType,
		Threshold: s.Threshold,
	}
}

var _ json.Marshaler = CategoricalDataDriftMetricThreshold{}

func (s CategoricalDataDriftMetricThreshold) MarshalJSON() ([]byte, error) {
	type wrapper CategoricalDataDriftMetricThreshold
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CategoricalDataDriftMetricThreshold: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CategoricalDataDriftMetricThreshold: %+v", err)
	}

	decoded["dataType"] = "Categorical"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CategoricalDataDriftMetricThreshold: %+v", err)
	}

	return encoded, nil
}
