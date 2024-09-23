package schedule

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataQualityMetricThresholdBase interface {
	DataQualityMetricThresholdBase() BaseDataQualityMetricThresholdBaseImpl
}

var _ DataQualityMetricThresholdBase = BaseDataQualityMetricThresholdBaseImpl{}

type BaseDataQualityMetricThresholdBaseImpl struct {
	DataType  MonitoringFeatureDataType `json:"dataType"`
	Threshold *MonitoringThreshold      `json:"threshold,omitempty"`
}

func (s BaseDataQualityMetricThresholdBaseImpl) DataQualityMetricThresholdBase() BaseDataQualityMetricThresholdBaseImpl {
	return s
}

var _ DataQualityMetricThresholdBase = RawDataQualityMetricThresholdBaseImpl{}

// RawDataQualityMetricThresholdBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDataQualityMetricThresholdBaseImpl struct {
	dataQualityMetricThresholdBase BaseDataQualityMetricThresholdBaseImpl
	Type                           string
	Values                         map[string]interface{}
}

func (s RawDataQualityMetricThresholdBaseImpl) DataQualityMetricThresholdBase() BaseDataQualityMetricThresholdBaseImpl {
	return s.dataQualityMetricThresholdBase
}

func UnmarshalDataQualityMetricThresholdBaseImplementation(input []byte) (DataQualityMetricThresholdBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DataQualityMetricThresholdBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["dataType"]; ok {
		value = fmt.Sprintf("%v", v)
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

	var parent BaseDataQualityMetricThresholdBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDataQualityMetricThresholdBaseImpl: %+v", err)
	}

	return RawDataQualityMetricThresholdBaseImpl{
		dataQualityMetricThresholdBase: parent,
		Type:                           value,
		Values:                         temp,
	}, nil

}
