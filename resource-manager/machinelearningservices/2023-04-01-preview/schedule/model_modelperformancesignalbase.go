package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MonitoringSignalBase = ModelPerformanceSignalBase{}

type ModelPerformanceSignalBase struct {
	BaselineData    MonitoringInputData                 `json:"baselineData"`
	DataSegment     *MonitoringDataSegment              `json:"dataSegment,omitempty"`
	MetricThreshold ModelPerformanceMetricThresholdBase `json:"metricThreshold"`
	TargetData      MonitoringInputData                 `json:"targetData"`

	// Fields inherited from MonitoringSignalBase
	LookbackPeriod *string                     `json:"lookbackPeriod,omitempty"`
	Mode           *MonitoringNotificationMode `json:"mode,omitempty"`
}

var _ json.Marshaler = ModelPerformanceSignalBase{}

func (s ModelPerformanceSignalBase) MarshalJSON() ([]byte, error) {
	type wrapper ModelPerformanceSignalBase
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ModelPerformanceSignalBase: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ModelPerformanceSignalBase: %+v", err)
	}
	decoded["signalType"] = "ModelPerformance"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ModelPerformanceSignalBase: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ModelPerformanceSignalBase{}

func (s *ModelPerformanceSignalBase) UnmarshalJSON(bytes []byte) error {
	type alias ModelPerformanceSignalBase
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into ModelPerformanceSignalBase: %+v", err)
	}

	s.BaselineData = decoded.BaselineData
	s.DataSegment = decoded.DataSegment
	s.LookbackPeriod = decoded.LookbackPeriod
	s.Mode = decoded.Mode
	s.TargetData = decoded.TargetData

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ModelPerformanceSignalBase into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["metricThreshold"]; ok {
		impl, err := unmarshalModelPerformanceMetricThresholdBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'MetricThreshold' for 'ModelPerformanceSignalBase': %+v", err)
		}
		s.MetricThreshold = impl
	}
	return nil
}
