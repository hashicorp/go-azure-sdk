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

// RawModeOfTransitImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDataDriftMetricThresholdBaseImpl struct {
	Type   string
	Values map[string]interface{}
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

	out := RawDataDriftMetricThresholdBaseImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
