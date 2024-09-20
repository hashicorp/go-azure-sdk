package schedule

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MonitoringInputDataBase interface {
	MonitoringInputDataBase() BaseMonitoringInputDataBaseImpl
}

var _ MonitoringInputDataBase = BaseMonitoringInputDataBaseImpl{}

type BaseMonitoringInputDataBaseImpl struct {
	Columns       *map[string]string      `json:"columns,omitempty"`
	DataContext   *string                 `json:"dataContext,omitempty"`
	InputDataType MonitoringInputDataType `json:"inputDataType"`
	JobInputType  JobInputType            `json:"jobInputType"`
	Uri           string                  `json:"uri"`
}

func (s BaseMonitoringInputDataBaseImpl) MonitoringInputDataBase() BaseMonitoringInputDataBaseImpl {
	return s
}

var _ MonitoringInputDataBase = RawMonitoringInputDataBaseImpl{}

// RawMonitoringInputDataBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawMonitoringInputDataBaseImpl struct {
	monitoringInputDataBase BaseMonitoringInputDataBaseImpl
	Type                    string
	Values                  map[string]interface{}
}

func (s RawMonitoringInputDataBaseImpl) MonitoringInputDataBase() BaseMonitoringInputDataBaseImpl {
	return s.monitoringInputDataBase
}

func UnmarshalMonitoringInputDataBaseImplementation(input []byte) (MonitoringInputDataBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling MonitoringInputDataBase into map[string]interface: %+v", err)
	}

	value, ok := temp["inputDataType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Fixed") {
		var out FixedInputData
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FixedInputData: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Rolling") {
		var out RollingInputData
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RollingInputData: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Static") {
		var out StaticInputData
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into StaticInputData: %+v", err)
		}
		return out, nil
	}

	var parent BaseMonitoringInputDataBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseMonitoringInputDataBaseImpl: %+v", err)
	}

	return RawMonitoringInputDataBaseImpl{
		monitoringInputDataBase: parent,
		Type:                    value,
		Values:                  temp,
	}, nil

}
