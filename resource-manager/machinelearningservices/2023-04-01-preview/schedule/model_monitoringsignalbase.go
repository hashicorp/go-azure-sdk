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

	if strings.EqualFold(value, "ModelPerformance") {
		var out ModelPerformanceSignalBase
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ModelPerformanceSignalBase: %+v", err)
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

	type RawMonitoringSignalBaseImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawMonitoringSignalBaseImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
