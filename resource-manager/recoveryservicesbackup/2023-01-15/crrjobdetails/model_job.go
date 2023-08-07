package crrjobdetails

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Job interface {
}

// RawModeOfTransitImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawJobImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalJobImplementation(input []byte) (Job, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Job into map[string]interface: %+v", err)
	}

	value, ok := temp["jobType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "AzureIaaSVMJob") {
		var out AzureIaaSVMJob
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureIaaSVMJob: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureStorageJob") {
		var out AzureStorageJob
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureStorageJob: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AzureWorkloadJob") {
		var out AzureWorkloadJob
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AzureWorkloadJob: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "DpmJob") {
		var out DpmJob
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DpmJob: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "MabJob") {
		var out MabJob
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into MabJob: %+v", err)
		}
		return out, nil
	}

	out := RawJobImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
