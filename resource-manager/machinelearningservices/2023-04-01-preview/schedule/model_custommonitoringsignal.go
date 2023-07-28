package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MonitoringSignalBase = CustomMonitoringSignal{}

type CustomMonitoringSignal struct {
	ComponentId      string                          `json:"componentId"`
	InputAssets      *map[string]MonitoringInputData `json:"inputAssets,omitempty"`
	MetricThresholds []CustomMetricThreshold         `json:"metricThresholds"`

	// Fields inherited from MonitoringSignalBase
	LookbackPeriod *string                     `json:"lookbackPeriod,omitempty"`
	Mode           *MonitoringNotificationMode `json:"mode,omitempty"`
}

var _ json.Marshaler = CustomMonitoringSignal{}

func (s CustomMonitoringSignal) MarshalJSON() ([]byte, error) {
	type wrapper CustomMonitoringSignal
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling CustomMonitoringSignal: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomMonitoringSignal: %+v", err)
	}
	decoded["signalType"] = "Custom"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CustomMonitoringSignal: %+v", err)
	}

	return encoded, nil
}
