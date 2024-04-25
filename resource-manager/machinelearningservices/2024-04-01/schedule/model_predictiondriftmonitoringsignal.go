package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MonitoringSignalBase = PredictionDriftMonitoringSignal{}

type PredictionDriftMonitoringSignal struct {
	FeatureDataTypeOverride *map[string]MonitoringFeatureDataType `json:"featureDataTypeOverride,omitempty"`
	MetricThresholds        []PredictionDriftMetricThresholdBase  `json:"metricThresholds"`
	ProductionData          MonitoringInputDataBase               `json:"productionData"`
	ReferenceData           MonitoringInputDataBase               `json:"referenceData"`

	// Fields inherited from MonitoringSignalBase
	NotificationTypes *[]MonitoringNotificationType `json:"notificationTypes,omitempty"`
	Properties        *map[string]string            `json:"properties,omitempty"`
}

var _ json.Marshaler = PredictionDriftMonitoringSignal{}

func (s PredictionDriftMonitoringSignal) MarshalJSON() ([]byte, error) {
	type wrapper PredictionDriftMonitoringSignal
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PredictionDriftMonitoringSignal: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PredictionDriftMonitoringSignal: %+v", err)
	}
	decoded["signalType"] = "PredictionDrift"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PredictionDriftMonitoringSignal: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &PredictionDriftMonitoringSignal{}

func (s *PredictionDriftMonitoringSignal) UnmarshalJSON(bytes []byte) error {
	type alias PredictionDriftMonitoringSignal
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into PredictionDriftMonitoringSignal: %+v", err)
	}

	s.FeatureDataTypeOverride = decoded.FeatureDataTypeOverride
	s.NotificationTypes = decoded.NotificationTypes
	s.Properties = decoded.Properties

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling PredictionDriftMonitoringSignal into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["metricThresholds"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling MetricThresholds into list []json.RawMessage: %+v", err)
		}

		output := make([]PredictionDriftMetricThresholdBase, 0)
		for i, val := range listTemp {
			impl, err := unmarshalPredictionDriftMetricThresholdBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'MetricThresholds' for 'PredictionDriftMonitoringSignal': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.MetricThresholds = output
	}

	if v, ok := temp["productionData"]; ok {
		impl, err := unmarshalMonitoringInputDataBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ProductionData' for 'PredictionDriftMonitoringSignal': %+v", err)
		}
		s.ProductionData = impl
	}

	if v, ok := temp["referenceData"]; ok {
		impl, err := unmarshalMonitoringInputDataBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ReferenceData' for 'PredictionDriftMonitoringSignal': %+v", err)
		}
		s.ReferenceData = impl
	}
	return nil
}
