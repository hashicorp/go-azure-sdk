package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataDriftMetricThresholdBase = NumericalDataDriftMetricThreshold{}

type NumericalDataDriftMetricThreshold struct {
	Metric NumericalDataDriftMetric `json:"metric"`

	// Fields inherited from DataDriftMetricThresholdBase
	Threshold *MonitoringThreshold `json:"threshold,omitempty"`
}

var _ json.Marshaler = NumericalDataDriftMetricThreshold{}

func (s NumericalDataDriftMetricThreshold) MarshalJSON() ([]byte, error) {
	type wrapper NumericalDataDriftMetricThreshold
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling NumericalDataDriftMetricThreshold: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling NumericalDataDriftMetricThreshold: %+v", err)
	}
	decoded["dataType"] = "Numerical"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling NumericalDataDriftMetricThreshold: %+v", err)
	}

	return encoded, nil
}
