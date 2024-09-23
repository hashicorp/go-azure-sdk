package replicationjobs

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobDetails interface {
	JobDetails() BaseJobDetailsImpl
}

var _ JobDetails = BaseJobDetailsImpl{}

type BaseJobDetailsImpl struct {
	AffectedObjectDetails *map[string]string `json:"affectedObjectDetails,omitempty"`
	InstanceType          string             `json:"instanceType"`
}

func (s BaseJobDetailsImpl) JobDetails() BaseJobDetailsImpl {
	return s
}

var _ JobDetails = RawJobDetailsImpl{}

// RawJobDetailsImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawJobDetailsImpl struct {
	jobDetails BaseJobDetailsImpl
	Type       string
	Values     map[string]interface{}
}

func (s RawJobDetailsImpl) JobDetails() BaseJobDetailsImpl {
	return s.jobDetails
}

func UnmarshalJobDetailsImplementation(input []byte) (JobDetails, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling JobDetails into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["instanceType"]; ok {
		value = fmt.Sprintf("%v", v)
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

	var parent BaseJobDetailsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseJobDetailsImpl: %+v", err)
	}

	return RawJobDetailsImpl{
		jobDetails: parent,
		Type:       value,
		Values:     temp,
	}, nil

}
