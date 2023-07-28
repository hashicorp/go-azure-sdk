package job

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SparkJobEntry interface {
}

func unmarshalSparkJobEntryImplementation(input []byte) (SparkJobEntry, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling SparkJobEntry into map[string]interface: %+v", err)
	}

	value, ok := temp["sparkJobEntryType"].(string)
	if !ok {
		return nil, nil
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

	type RawSparkJobEntryImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawSparkJobEntryImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
