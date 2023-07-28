package schedule

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MonitoringFeatureFilterBase interface {
}

func unmarshalMonitoringFeatureFilterBaseImplementation(input []byte) (MonitoringFeatureFilterBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling MonitoringFeatureFilterBase into map[string]interface: %+v", err)
	}

	value, ok := temp["filterType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "AllFeatures") {
		var out AllFeatures
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AllFeatures: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "FeatureSubset") {
		var out FeatureSubset
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FeatureSubset: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "TopNByAttribution") {
		var out TopNFeaturesByAttribution
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TopNFeaturesByAttribution: %+v", err)
		}
		return out, nil
	}

	type RawMonitoringFeatureFilterBaseImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawMonitoringFeatureFilterBaseImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
