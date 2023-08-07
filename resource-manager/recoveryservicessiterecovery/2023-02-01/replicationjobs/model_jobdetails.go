package replicationjobs

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobDetails interface {
}

// RawModeOfTransitImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawJobDetailsImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalJobDetailsImplementation(input []byte) (JobDetails, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling JobDetails into map[string]interface: %+v", err)
	}

	value, ok := temp["instanceType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "AsrJobDetails") {
		var out AsrJobDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AsrJobDetails: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ExportJobDetails") {
		var out ExportJobDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ExportJobDetails: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "FailoverJobDetails") {
		var out FailoverJobDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into FailoverJobDetails: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "SwitchProtectionJobDetails") {
		var out SwitchProtectionJobDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SwitchProtectionJobDetails: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "TestFailoverJobDetails") {
		var out TestFailoverJobDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into TestFailoverJobDetails: %+v", err)
		}
		return out, nil
	}

	out := RawJobDetailsImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
