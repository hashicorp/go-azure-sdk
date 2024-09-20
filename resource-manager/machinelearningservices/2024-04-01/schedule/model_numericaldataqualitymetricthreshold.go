package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataQualityMetricThresholdBase = NumericalDataQualityMetricThreshold{}

type NumericalDataQualityMetricThreshold struct {
	Metric NumericalDataQualityMetric `json:"metric"`

	// Fields inherited from DataQualityMetricThresholdBase

	DataType  MonitoringFeatureDataType `json:"dataType"`
	Threshold *MonitoringThreshold      `json:"threshold,omitempty"`
}

func (s NumericalDataQualityMetricThreshold) DataQualityMetricThresholdBase() BaseDataQualityMetricThresholdBaseImpl {
	return BaseDataQualityMetricThresholdBaseImpl{
		DataType:  s.DataType,
		Threshold: s.Threshold,
	}
}

var _ json.Marshaler = NumericalDataQualityMetricThreshold{}

func (s NumericalDataQualityMetricThreshold) MarshalJSON() ([]byte, error) {
	type wrapper NumericalDataQualityMetricThreshold
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NumericalDataQualityMetricThreshold: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NumericalDataQualityMetricThreshold: %+v", err)
	}

	decoded["dataType"] = "Numerical"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NumericalDataQualityMetricThreshold: %+v", err)
	}

	return encoded, nil
}
