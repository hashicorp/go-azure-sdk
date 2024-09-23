package schedule

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MonitorComputeConfigurationBase interface {
	MonitorComputeConfigurationBase() BaseMonitorComputeConfigurationBaseImpl
}

var _ MonitorComputeConfigurationBase = BaseMonitorComputeConfigurationBaseImpl{}

type BaseMonitorComputeConfigurationBaseImpl struct {
	ComputeType MonitorComputeType `json:"computeType"`
}

func (s BaseMonitorComputeConfigurationBaseImpl) MonitorComputeConfigurationBase() BaseMonitorComputeConfigurationBaseImpl {
	return s
}

var _ MonitorComputeConfigurationBase = RawMonitorComputeConfigurationBaseImpl{}

// RawMonitorComputeConfigurationBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawMonitorComputeConfigurationBaseImpl struct {
	monitorComputeConfigurationBase BaseMonitorComputeConfigurationBaseImpl
	Type                            string
	Values                          map[string]interface{}
}

func (s RawMonitorComputeConfigurationBaseImpl) MonitorComputeConfigurationBase() BaseMonitorComputeConfigurationBaseImpl {
	return s.monitorComputeConfigurationBase
}

func UnmarshalMonitorComputeConfigurationBaseImplementation(input []byte) (MonitorComputeConfigurationBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling MonitorComputeConfigurationBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["computeType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "ServerlessSpark") {
		var out MonitorServerlessSparkCompute
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MonitorServerlessSparkCompute: %+v", err)
		}
		return out, nil
	}

	var parent BaseMonitorComputeConfigurationBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseMonitorComputeConfigurationBaseImpl: %+v", err)
	}

	return RawMonitorComputeConfigurationBaseImpl{
		monitorComputeConfigurationBase: parent,
		Type:                            value,
		Values:                          temp,
	}, nil

}
