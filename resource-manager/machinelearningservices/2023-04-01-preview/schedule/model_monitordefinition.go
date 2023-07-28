package schedule

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MonitorDefinition struct {
	AlertNotificationSetting MonitoringAlertNotificationSettingsBase `json:"alertNotificationSetting"`
	ComputeId                string                                  `json:"computeId"`
	MonitoringTarget         *string                                 `json:"monitoringTarget,omitempty"`
	Signals                  map[string]MonitoringSignalBase         `json:"signals"`
}

var _ json.Unmarshaler = &MonitorDefinition{}

func (s *MonitorDefinition) UnmarshalJSON(bytes []byte) error {
	type alias MonitorDefinition
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into MonitorDefinition: %+v", err)
	}

	s.ComputeId = decoded.ComputeId
	s.MonitoringTarget = decoded.MonitoringTarget

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling MonitorDefinition into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["alertNotificationSetting"]; ok {
		impl, err := unmarshalMonitoringAlertNotificationSettingsBaseImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'AlertNotificationSetting' for 'MonitorDefinition': %+v", err)
		}
		s.AlertNotificationSetting = impl
	}

	if v, ok := temp["signals"]; ok {
		var dictionaryTemp map[string]json.RawMessage
		if err := json.Unmarshal(v, &dictionaryTemp); err != nil {
			return fmt.Errorf("unmarshaling Signals into dictionary map[string]json.RawMessage: %+v", err)
		}

		output := make(map[string]MonitoringSignalBase)
		for key, val := range dictionaryTemp {
			impl, err := unmarshalMonitoringSignalBaseImplementation(val)
			if err != nil {
				return fmt.Errorf("unmarshaling key %q field 'Signals' for 'MonitorDefinition': %+v", key, err)
			}
			output[key] = impl
		}
		s.Signals = output
	}
	return nil
}
