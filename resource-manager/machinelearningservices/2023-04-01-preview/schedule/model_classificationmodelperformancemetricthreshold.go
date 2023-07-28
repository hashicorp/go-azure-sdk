package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ModelPerformanceMetricThresholdBase = ClassificationModelPerformanceMetricThreshold{}

type ClassificationModelPerformanceMetricThreshold struct {
	Metric ClassificationModelPerformanceMetric `json:"metric"`

	// Fields inherited from ModelPerformanceMetricThresholdBase
	Threshold *MonitoringThreshold `json:"threshold,omitempty"`
}

var _ json.Marshaler = ClassificationModelPerformanceMetricThreshold{}

func (s ClassificationModelPerformanceMetricThreshold) MarshalJSON() ([]byte, error) {
	type wrapper ClassificationModelPerformanceMetricThreshold
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ClassificationModelPerformanceMetricThreshold: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ClassificationModelPerformanceMetricThreshold: %+v", err)
	}
	decoded["modelType"] = "Classification"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ClassificationModelPerformanceMetricThreshold: %+v", err)
	}

	return encoded, nil
}
