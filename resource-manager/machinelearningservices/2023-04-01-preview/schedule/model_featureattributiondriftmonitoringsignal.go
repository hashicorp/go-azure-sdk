package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MonitoringSignalBase = FeatureAttributionDriftMonitoringSignal{}

type FeatureAttributionDriftMonitoringSignal struct {
	BaselineData    MonitoringInputData               `json:"baselineData"`
	MetricThreshold FeatureAttributionMetricThreshold `json:"metricThreshold"`
	ModelType       MonitoringModelType               `json:"modelType"`
	TargetData      MonitoringInputData               `json:"targetData"`

	// Fields inherited from MonitoringSignalBase
	LookbackPeriod *string                     `json:"lookbackPeriod,omitempty"`
	Mode           *MonitoringNotificationMode `json:"mode,omitempty"`
}

var _ json.Marshaler = FeatureAttributionDriftMonitoringSignal{}

func (s FeatureAttributionDriftMonitoringSignal) MarshalJSON() ([]byte, error) {
	type wrapper FeatureAttributionDriftMonitoringSignal
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling FeatureAttributionDriftMonitoringSignal: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling FeatureAttributionDriftMonitoringSignal: %+v", err)
	}
	decoded["signalType"] = "FeatureAttributionDrift"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling FeatureAttributionDriftMonitoringSignal: %+v", err)
	}

	return encoded, nil
}
