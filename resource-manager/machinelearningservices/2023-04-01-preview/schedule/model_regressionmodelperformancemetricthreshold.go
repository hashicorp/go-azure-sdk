package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ModelPerformanceMetricThresholdBase = RegressionModelPerformanceMetricThreshold{}

type RegressionModelPerformanceMetricThreshold struct {
	Metric RegressionModelPerformanceMetric `json:"metric"`

	// Fields inherited from ModelPerformanceMetricThresholdBase
	Threshold *MonitoringThreshold `json:"threshold,omitempty"`
}

var _ json.Marshaler = RegressionModelPerformanceMetricThreshold{}

func (s RegressionModelPerformanceMetricThreshold) MarshalJSON() ([]byte, error) {
	type wrapper RegressionModelPerformanceMetricThreshold
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RegressionModelPerformanceMetricThreshold: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RegressionModelPerformanceMetricThreshold: %+v", err)
	}
	decoded["modelType"] = "Regression"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RegressionModelPerformanceMetricThreshold: %+v", err)
	}

	return encoded, nil
}
