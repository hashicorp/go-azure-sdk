package schedule

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ForecastHorizon interface {
	ForecastHorizon() BaseForecastHorizonImpl
}

var _ ForecastHorizon = BaseForecastHorizonImpl{}

type BaseForecastHorizonImpl struct {
	Mode ForecastHorizonMode `json:"mode"`
}

func (s BaseForecastHorizonImpl) ForecastHorizon() BaseForecastHorizonImpl {
	return s
}

var _ ForecastHorizon = RawForecastHorizonImpl{}

// RawForecastHorizonImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawForecastHorizonImpl struct {
	forecastHorizon BaseForecastHorizonImpl
	Type            string
	Values          map[string]interface{}
}

func (s RawForecastHorizonImpl) ForecastHorizon() BaseForecastHorizonImpl {
	return s.forecastHorizon
}

func UnmarshalForecastHorizonImplementation(input []byte) (ForecastHorizon, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ForecastHorizon into map[string]interface: %+v", err)
	}

	value, ok := temp["mode"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Auto") {
		var out AutoForecastHorizon
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AutoForecastHorizon: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Custom") {
		var out CustomForecastHorizon
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CustomForecastHorizon: %+v", err)
		}
		return out, nil
	}

	var parent BaseForecastHorizonImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseForecastHorizonImpl: %+v", err)
	}

	return RawForecastHorizonImpl{
		forecastHorizon: parent,
		Type:            value,
		Values:          temp,
	}, nil

}
