package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ MonitoringSignalBase = CustomMonitoringSignal{}

type CustomMonitoringSignal struct {
	ComponentId      string                              `json:"componentId"`
	InputAssets      *map[string]MonitoringInputDataBase `json:"inputAssets,omitempty"`
	Inputs           *map[string]JobInput                `json:"inputs,omitempty"`
	MetricThresholds []CustomMetricThreshold             `json:"metricThresholds"`

	// Fields inherited from MonitoringSignalBase

	NotificationTypes *[]MonitoringNotificationType `json:"notificationTypes,omitempty"`
	Properties        *map[string]string            `json:"properties,omitempty"`
	SignalType        MonitoringSignalType          `json:"signalType"`
}

func (s CustomMonitoringSignal) MonitoringSignalBase() BaseMonitoringSignalBaseImpl {
	return BaseMonitoringSignalBaseImpl{
		NotificationTypes: s.NotificationTypes,
		Properties:        s.Properties,
		SignalType:        s.SignalType,
	}
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
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling CustomMonitoringSignal: %+v", err)
	}

	decoded["signalType"] = "Custom"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling CustomMonitoringSignal: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &CustomMonitoringSignal{}

func (s *CustomMonitoringSignal) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ComponentId       string                        `json:"componentId"`
		MetricThresholds  []CustomMetricThreshold       `json:"metricThresholds"`
		NotificationTypes *[]MonitoringNotificationType `json:"notificationTypes,omitempty"`
		Properties        *map[string]string            `json:"properties,omitempty"`
		SignalType        MonitoringSignalType          `json:"signalType"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ComponentId = decoded.ComponentId
	s.MetricThresholds = decoded.MetricThresholds
	s.NotificationTypes = decoded.NotificationTypes
	s.Properties = decoded.Properties
	s.SignalType = decoded.SignalType

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling CustomMonitoringSignal into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["inputAssets"]; ok {
		var dictionaryTemp map[string]json.RawMessage
		if err := json.Unmarshal(v, &dictionaryTemp); err != nil {
			return fmt.Errorf("unmarshaling InputAssets into dictionary map[string]json.RawMessage: %+v", err)
		}

		output := make(map[string]MonitoringInputDataBase)
		for key, val := range dictionaryTemp {
			impl, err := UnmarshalMonitoringInputDataBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling key %q field 'InputAssets' for 'CustomMonitoringSignal': %+v", key, err)
			}
			output[key] = impl
		}
		s.InputAssets = &output
	}

	if v, ok := temp["inputs"]; ok {
		var dictionaryTemp map[string]json.RawMessage
		if err := json.Unmarshal(v, &dictionaryTemp); err != nil {
			return fmt.Errorf("unmarshaling Inputs into dictionary map[string]json.RawMessage: %+v", err)
		}

		output := make(map[string]JobInput)
		for key, val := range dictionaryTemp {
			impl, err := UnmarshalJobInputImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling key %q field 'Inputs' for 'CustomMonitoringSignal': %+v", key, err)
			}
			output[key] = impl
		}
		s.Inputs = &output
	}

	return nil
}
