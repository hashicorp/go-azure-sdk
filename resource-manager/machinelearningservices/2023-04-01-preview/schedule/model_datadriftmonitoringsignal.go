package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MonitoringSignalBase = DataDriftMonitoringSignal{}

type DataDriftMonitoringSignal struct {
	BaselineData     MonitoringInputData            `json:"baselineData"`
	DataSegment      *MonitoringDataSegment         `json:"dataSegment,omitempty"`
	Features         MonitoringFeatureFilterBase    `json:"features"`
	MetricThresholds []DataDriftMetricThresholdBase `json:"metricThresholds"`
	TargetData       MonitoringInputData            `json:"targetData"`

	// Fields inherited from MonitoringSignalBase
	LookbackPeriod *string                     `json:"lookbackPeriod,omitempty"`
	Mode           *MonitoringNotificationMode `json:"mode,omitempty"`
}

var _ json.Marshaler = DataDriftMonitoringSignal{}

func (s DataDriftMonitoringSignal) MarshalJSON() ([]byte, error) {
	type wrapper DataDriftMonitoringSignal
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DataDriftMonitoringSignal: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DataDriftMonitoringSignal: %+v", err)
	}
	decoded["signalType"] = "DataDrift"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DataDriftMonitoringSignal: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &DataDriftMonitoringSignal{}

func (s *DataDriftMonitoringSignal) UnmarshalJSON(bytes []byte) error {
	type alias DataDriftMonitoringSignal
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into DataDriftMonitoringSignal: %+v", err)
	}

	s.BaselineData = decoded.BaselineData
	s.DataSegment = decoded.DataSegment
	s.LookbackPeriod = decoded.LookbackPeriod
	s.Mode = decoded.Mode
	s.TargetData = decoded.TargetData

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DataDriftMonitoringSignal into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["features"]; ok {
		impl, err := unmarshalMonitoringFeatureFilterBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Features' for 'DataDriftMonitoringSignal': %+v", err)
		}
		s.Features = impl
	}

	if v, ok := temp["metricThresholds"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling MetricThresholds into list []json.RawMessage: %+v", err)
		}

		output := make([]DataDriftMetricThresholdBase, 0)
		for i, val := range listTemp {
			impl, err := unmarshalDataDriftMetricThresholdBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'MetricThresholds' for 'DataDriftMonitoringSignal': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.MetricThresholds = output
	}
	return nil
}
