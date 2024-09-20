package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MonitoringSignalBase = FeatureAttributionDriftMonitoringSignal{}

type FeatureAttributionDriftMonitoringSignal struct {
	FeatureDataTypeOverride   *map[string]MonitoringFeatureDataType `json:"featureDataTypeOverride,omitempty"`
	FeatureImportanceSettings FeatureImportanceSettings             `json:"featureImportanceSettings"`
	MetricThreshold           FeatureAttributionMetricThreshold     `json:"metricThreshold"`
	ProductionData            []MonitoringInputDataBase             `json:"productionData"`
	ReferenceData             MonitoringInputDataBase               `json:"referenceData"`

	// Fields inherited from MonitoringSignalBase

	NotificationTypes *[]MonitoringNotificationType `json:"notificationTypes,omitempty"`
	Properties        *map[string]string            `json:"properties,omitempty"`
	SignalType        MonitoringSignalType          `json:"signalType"`
}

func (s FeatureAttributionDriftMonitoringSignal) MonitoringSignalBase() BaseMonitoringSignalBaseImpl {
	return BaseMonitoringSignalBaseImpl{
		NotificationTypes: s.NotificationTypes,
		Properties:        s.Properties,
		SignalType:        s.SignalType,
	}
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
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling FeatureAttributionDriftMonitoringSignal: %+v", err)
	}

	decoded["signalType"] = "FeatureAttributionDrift"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling FeatureAttributionDriftMonitoringSignal: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &FeatureAttributionDriftMonitoringSignal{}

func (s *FeatureAttributionDriftMonitoringSignal) UnmarshalJSON(bytes []byte) error {
	type alias FeatureAttributionDriftMonitoringSignal
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into FeatureAttributionDriftMonitoringSignal: %+v", err)
	}

	s.FeatureDataTypeOverride = decoded.FeatureDataTypeOverride
	s.FeatureImportanceSettings = decoded.FeatureImportanceSettings
	s.MetricThreshold = decoded.MetricThreshold
	s.NotificationTypes = decoded.NotificationTypes
	s.Properties = decoded.Properties
	s.SignalType = decoded.SignalType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling FeatureAttributionDriftMonitoringSignal into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["productionData"]; ok {
		var listTemp []json.RawMessage
		if err := json.Unmarshal(v, &listTemp); err != nil {
			return fmt.Errorf("unmarshaling ProductionData into list []json.RawMessage: %+v", err)
		}

		output := make([]MonitoringInputDataBase, 0)
		for i, val := range listTemp {
			impl, err := UnmarshalMonitoringInputDataBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling index %d field 'ProductionData' for 'FeatureAttributionDriftMonitoringSignal': %+v", i, err)
			}
			output = append(output, impl)
		}
		s.ProductionData = output
	}

	if v, ok := temp["referenceData"]; ok {
		impl, err := UnmarshalMonitoringInputDataBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ReferenceData' for 'FeatureAttributionDriftMonitoringSignal': %+v", err)
		}
		s.ReferenceData = impl
	}
	return nil
}
