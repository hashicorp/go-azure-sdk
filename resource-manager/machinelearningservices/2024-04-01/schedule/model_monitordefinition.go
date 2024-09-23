package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MonitorDefinition struct {
	AlertNotificationSettings *MonitorNotificationSettings    `json:"alertNotificationSettings,omitempty"`
	ComputeConfiguration      MonitorComputeConfigurationBase `json:"computeConfiguration"`
	MonitoringTarget          *MonitoringTarget               `json:"monitoringTarget,omitempty"`
	Signals                   map[string]MonitoringSignalBase `json:"signals"`
}

var _ json.Unmarshaler = &MonitorDefinition{}

func (s *MonitorDefinition) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		AlertNotificationSettings *MonitorNotificationSettings `json:"alertNotificationSettings,omitempty"`
		MonitoringTarget          *MonitoringTarget            `json:"monitoringTarget,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.AlertNotificationSettings = decoded.AlertNotificationSettings
	s.MonitoringTarget = decoded.MonitoringTarget

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling MonitorDefinition into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["computeConfiguration"]; ok {
		impl, err := UnmarshalMonitorComputeConfigurationBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ComputeConfiguration' for 'MonitorDefinition': %+v", err)
		}
		s.ComputeConfiguration = impl
	}

	if v, ok := temp["signals"]; ok {
		var dictionaryTemp map[string]json.RawMessage
		if err := json.Unmarshal(v, &dictionaryTemp); err != nil {
			return fmt.Errorf("unmarshaling Signals into dictionary map[string]json.RawMessage: %+v", err)
		}

		output := make(map[string]MonitoringSignalBase)
		for key, val := range dictionaryTemp {
			impl, err := UnmarshalMonitoringSignalBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling key %q field 'Signals' for 'MonitorDefinition': %+v", key, err)
			}
			output[key] = impl
		}
		s.Signals = output
	}

	return nil
}
