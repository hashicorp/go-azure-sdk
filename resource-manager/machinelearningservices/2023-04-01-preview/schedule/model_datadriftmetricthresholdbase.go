package schedule

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataDriftMetricThresholdBase interface {
}

func unmarshalDataDriftMetricThresholdBaseImplementation(input []byte) (DataDriftMetricThresholdBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DataDriftMetricThresholdBase into map[string]interface: %+v", err)
	}

	value, ok := temp["dataType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Categorical") {
		var out CategoricalDataDriftMetricThreshold
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CategoricalDataDriftMetricThreshold: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Numerical") {
		var out NumericalDataDriftMetricThreshold
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NumericalDataDriftMetricThreshold: %+v", err)
		}
		return out, nil
	}

	type RawDataDriftMetricThresholdBaseImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawDataDriftMetricThresholdBaseImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
