package schedule

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataDriftMetricThresholdBase interface {
	DataDriftMetricThresholdBase() BaseDataDriftMetricThresholdBaseImpl
}

var _ DataDriftMetricThresholdBase = BaseDataDriftMetricThresholdBaseImpl{}

type BaseDataDriftMetricThresholdBaseImpl struct {
	DataType  MonitoringFeatureDataType `json:"dataType"`
	Threshold *MonitoringThreshold      `json:"threshold,omitempty"`
}

func (s BaseDataDriftMetricThresholdBaseImpl) DataDriftMetricThresholdBase() BaseDataDriftMetricThresholdBaseImpl {
	return s
}

var _ DataDriftMetricThresholdBase = RawDataDriftMetricThresholdBaseImpl{}

// RawDataDriftMetricThresholdBaseImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawDataDriftMetricThresholdBaseImpl struct {
	dataDriftMetricThresholdBase BaseDataDriftMetricThresholdBaseImpl
	Type                         string
	Values                       map[string]interface{}
}

func (s RawDataDriftMetricThresholdBaseImpl) DataDriftMetricThresholdBase() BaseDataDriftMetricThresholdBaseImpl {
	return s.dataDriftMetricThresholdBase
}

func (s RawDataDriftMetricThresholdBaseImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalDataDriftMetricThresholdBaseImplementation(input []byte) (DataDriftMetricThresholdBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DataDriftMetricThresholdBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["dataType"]; ok {
		value = fmt.Sprintf("%v", v)
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

	var parent BaseDataDriftMetricThresholdBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDataDriftMetricThresholdBaseImpl: %+v", err)
	}

	return RawDataDriftMetricThresholdBaseImpl{
		dataDriftMetricThresholdBase: parent,
		Type:                         value,
		Values:                       temp,
	}, nil

}
