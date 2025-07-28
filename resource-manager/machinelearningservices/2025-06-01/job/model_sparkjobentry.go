package job

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SparkJobEntry interface {
	SparkJobEntry() BaseSparkJobEntryImpl
}

var _ SparkJobEntry = BaseSparkJobEntryImpl{}

type BaseSparkJobEntryImpl struct {
	SparkJobEntryType SparkJobEntryType `json:"sparkJobEntryType"`
}

func (s BaseSparkJobEntryImpl) SparkJobEntry() BaseSparkJobEntryImpl {
	return s
}

var _ SparkJobEntry = RawSparkJobEntryImpl{}

// RawSparkJobEntryImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawSparkJobEntryImpl struct {
	sparkJobEntry BaseSparkJobEntryImpl
	Type          string
	Values        map[string]interface{}
}

func (s RawSparkJobEntryImpl) SparkJobEntry() BaseSparkJobEntryImpl {
	return s.sparkJobEntry
}

func UnmarshalSparkJobEntryImplementation(input []byte) (SparkJobEntry, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SparkJobEntry into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["sparkJobEntryType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "SparkJobPythonEntry") {
		var out SparkJobPythonEntry
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SparkJobPythonEntry: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SparkJobScalaEntry") {
		var out SparkJobScalaEntry
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SparkJobScalaEntry: %+v", err)
		}
		return out, nil
	}

	var parent BaseSparkJobEntryImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseSparkJobEntryImpl: %+v", err)
	}

	return RawSparkJobEntryImpl{
		sparkJobEntry: parent,
		Type:          value,
		Values:        temp,
	}, nil

}
