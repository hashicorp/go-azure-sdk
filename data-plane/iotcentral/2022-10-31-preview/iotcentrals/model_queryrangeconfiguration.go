package iotcentrals

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type QueryRangeConfiguration interface {
	QueryRangeConfiguration() BaseQueryRangeConfigurationImpl
}

var _ QueryRangeConfiguration = BaseQueryRangeConfigurationImpl{}

type BaseQueryRangeConfigurationImpl struct {
	Type string `json:"type"`
}

func (s BaseQueryRangeConfigurationImpl) QueryRangeConfiguration() BaseQueryRangeConfigurationImpl {
	return s
}

var _ QueryRangeConfiguration = RawQueryRangeConfigurationImpl{}

// RawQueryRangeConfigurationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawQueryRangeConfigurationImpl struct {
	queryRangeConfiguration BaseQueryRangeConfigurationImpl
	Type                    string
	Values                  map[string]interface{}
}

func (s RawQueryRangeConfigurationImpl) QueryRangeConfiguration() BaseQueryRangeConfigurationImpl {
	return s.queryRangeConfiguration
}

func UnmarshalQueryRangeConfigurationImplementation(input []byte) (QueryRangeConfiguration, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling QueryRangeConfiguration into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "count") {
		var out CountQueryRangeConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CountQueryRangeConfiguration: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "time") {
		var out TimeQueryRangeConfiguration
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TimeQueryRangeConfiguration: %+v", err)
		}
		return out, nil
	}

	var parent BaseQueryRangeConfigurationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseQueryRangeConfigurationImpl: %+v", err)
	}

	return RawQueryRangeConfigurationImpl{
		queryRangeConfiguration: parent,
		Type:                    value,
		Values:                  temp,
	}, nil

}
