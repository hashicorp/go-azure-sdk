package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataQualityMetricThresholdBase = CategoricalDataQualityMetricThreshold{}

type CategoricalDataQualityMetricThreshold struct {
	Metric CategoricalDataQualityMetric `json:"metric"`

	// Fields inherited from DataQualityMetricThresholdBase

	DataType  MonitoringFeatureDataType `json:"dataType"`
	Threshold *MonitoringThreshold      `json:"threshold,omitempty"`
}

func (s CategoricalDataQualityMetricThreshold) DataQualityMetricThresholdBase() BaseDataQualityMetricThresholdBaseImpl {
	return BaseDataQualityMetricThresholdBaseImpl{
		DataType:  s.DataType,
		Threshold: s.Threshold,
	}
}

var _ json.Marshaler = CategoricalDataQualityMetricThreshold{}

func (s CategoricalDataQualityMetricThreshold) MarshalJSON() ([]byte, error) {
	type wrapper CategoricalDataQualityMetricThreshold
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CategoricalDataQualityMetricThreshold: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CategoricalDataQualityMetricThreshold: %+v", err)
	}

	decoded["dataType"] = "Categorical"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CategoricalDataQualityMetricThreshold: %+v", err)
	}

	return encoded, nil
}
