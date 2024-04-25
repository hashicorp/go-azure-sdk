package schedule

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MonitoringSignalBase interface {
}

// RawMonitoringSignalBaseImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawMonitoringSignalBaseImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalMonitoringSignalBaseImplementation(input []byte) (MonitoringSignalBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling MonitoringSignalBase into map[string]interface: %+v", err)
	}

	value, ok := temp["signalType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Custom") {
		var out CustomMonitoringSignal
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomMonitoringSignal: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "DataDrift") {
		var out DataDriftMonitoringSignal
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DataDriftMonitoringSignal: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "DataQuality") {
		var out DataQualityMonitoringSignal
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DataQualityMonitoringSignal: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "FeatureAttributionDrift") {
		var out FeatureAttributionDriftMonitoringSignal
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FeatureAttributionDriftMonitoringSignal: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "PredictionDrift") {
		var out PredictionDriftMonitoringSignal
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into PredictionDriftMonitoringSignal: %+v", err)
		}
		return out, nil
	}

	out := RawMonitoringSignalBaseImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
