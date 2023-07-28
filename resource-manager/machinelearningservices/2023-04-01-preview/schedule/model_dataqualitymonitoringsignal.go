package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MonitoringSignalBase = DataQualityMonitoringSignal{}

type DataQualityMonitoringSignal struct {
	BaselineData     MonitoringInputData              `json:"baselineData"`
	Features         MonitoringFeatureFilterBase      `json:"features"`
	MetricThresholds []DataQualityMetricThresholdBase `json:"metricThresholds"`
	TargetData       MonitoringInputData              `json:"targetData"`

	// Fields inherited from MonitoringSignalBase
	LookbackPeriod *string                     `json:"lookbackPeriod,omitempty"`
	Mode           *MonitoringNotificationMode `json:"mode,omitempty"`
}

var _ json.Marshaler = DataQualityMonitoringSignal{}

func (s DataQualityMonitoringSignal) MarshalJSON() ([]byte, error) {
	type wrapper DataQualityMonitoringSignal
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DataQualityMonitoringSignal: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DataQualityMonitoringSignal: %+v", err)
	}
	decoded["signalType"] = "DataQuality"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DataQualityMonitoringSignal: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DataQualityMonitoringSignal{}

func (s *DataQualityMonitoringSignal) UnmarshalJSON(bytes []byte) error {
	type alias DataQualityMonitoringSignal
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into DataQualityMonitoringSignal: %+v", err)
	}

	s.BaselineData = decoded.BaselineData
	s.LookbackPeriod = decoded.LookbackPeriod
	s.Mode = decoded.Mode
	s.TargetData = decoded.TargetData

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DataQualityMonitoringSignal into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["features"]; ok {
		impl, err := unmarshalMonitoringFeatureFilterBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Features' for 'DataQualityMonitoringSignal': %+v", err)
		}
		s.Features = impl
	}

	if v, ok := temp["metricThresholds"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling MetricThresholds into list []json.RawMessage: %+v", err)
		}

		output := make([]DataQualityMetricThresholdBase, 0)
		for i, val := range listTemp {
			impl, err := unmarshalDataQualityMetricThresholdBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'MetricThresholds' for 'DataQualityMonitoringSignal': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.MetricThresholds = output
	}
	return nil
}
