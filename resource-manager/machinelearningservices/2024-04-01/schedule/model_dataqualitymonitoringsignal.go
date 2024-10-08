package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MonitoringSignalBase = DataQualityMonitoringSignal{}

type DataQualityMonitoringSignal struct {
	FeatureDataTypeOverride   *map[string]MonitoringFeatureDataType `json:"featureDataTypeOverride,omitempty"`
	FeatureImportanceSettings *FeatureImportanceSettings            `json:"featureImportanceSettings,omitempty"`
	Features                  MonitoringFeatureFilterBase           `json:"features"`
	MetricThresholds          []DataQualityMetricThresholdBase      `json:"metricThresholds"`
	ProductionData            MonitoringInputDataBase               `json:"productionData"`
	ReferenceData             MonitoringInputDataBase               `json:"referenceData"`

	// Fields inherited from MonitoringSignalBase

	NotificationTypes *[]MonitoringNotificationType `json:"notificationTypes,omitempty"`
	Properties        *map[string]string            `json:"properties,omitempty"`
	SignalType        MonitoringSignalType          `json:"signalType"`
}

func (s DataQualityMonitoringSignal) MonitoringSignalBase() BaseMonitoringSignalBaseImpl {
	return BaseMonitoringSignalBaseImpl{
		NotificationTypes: s.NotificationTypes,
		Properties:        s.Properties,
		SignalType:        s.SignalType,
	}
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
	if err = json.Unmarshal(encoded, &decoded); err != nil {
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
	var decoded struct {
		FeatureDataTypeOverride   *map[string]MonitoringFeatureDataType `json:"featureDataTypeOverride,omitempty"`
		FeatureImportanceSettings *FeatureImportanceSettings            `json:"featureImportanceSettings,omitempty"`
		NotificationTypes         *[]MonitoringNotificationType         `json:"notificationTypes,omitempty"`
		Properties                *map[string]string                    `json:"properties,omitempty"`
		SignalType                MonitoringSignalType                  `json:"signalType"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.FeatureDataTypeOverride = decoded.FeatureDataTypeOverride
	s.FeatureImportanceSettings = decoded.FeatureImportanceSettings
	s.NotificationTypes = decoded.NotificationTypes
	s.Properties = decoded.Properties
	s.SignalType = decoded.SignalType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling DataQualityMonitoringSignal into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["features"]; ok {
		impl, err := UnmarshalMonitoringFeatureFilterBaseImplementation(v)
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
			impl, err := UnmarshalDataQualityMetricThresholdBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'MetricThresholds' for 'DataQualityMonitoringSignal': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.MetricThresholds = output
	}

	if v, ok := temp["productionData"]; ok {
		impl, err := UnmarshalMonitoringInputDataBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ProductionData' for 'DataQualityMonitoringSignal': %+v", err)
		}
		s.ProductionData = impl
	}

	if v, ok := temp["referenceData"]; ok {
		impl, err := UnmarshalMonitoringInputDataBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ReferenceData' for 'DataQualityMonitoringSignal': %+v", err)
		}
		s.ReferenceData = impl
	}

	return nil
}
