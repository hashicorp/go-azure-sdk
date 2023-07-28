package schedule

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataQualityMetricThresholdBase interface {
}

func unmarshalDataQualityMetricThresholdBaseImplementation(input []byte) (DataQualityMetricThresholdBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DataQualityMetricThresholdBase into map[string]interface: %+v", err)
	}

	value, ok := temp["dataType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Categorical") {
		var out CategoricalDataQualityMetricThreshold
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CategoricalDataQualityMetricThreshold: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Numerical") {
		var out NumericalDataQualityMetricThreshold
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into NumericalDataQualityMetricThreshold: %+v", err)
		}
		return out, nil
	}

	type RawDataQualityMetricThresholdBaseImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawDataQualityMetricThresholdBaseImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
