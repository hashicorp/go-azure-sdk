package schedule

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MonitoringInputDataBase interface {
}

// RawMonitoringInputDataBaseImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawMonitoringInputDataBaseImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalMonitoringInputDataBaseImplementation(input []byte) (MonitoringInputDataBase, error) {
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

	out := RawMonitoringInputDataBaseImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
