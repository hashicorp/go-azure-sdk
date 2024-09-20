package schedule

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobLimits interface {
	JobLimits() BaseJobLimitsImpl
}

var _ JobLimits = BaseJobLimitsImpl{}

type BaseJobLimitsImpl struct {
	JobLimitsType JobLimitsType `json:"jobLimitsType"`
	Timeout       *string       `json:"timeout,omitempty"`
}

func (s BaseJobLimitsImpl) JobLimits() BaseJobLimitsImpl {
	return s
}

var _ JobLimits = RawJobLimitsImpl{}

// RawJobLimitsImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawJobLimitsImpl struct {
	jobLimits BaseJobLimitsImpl
	Type      string
	Values    map[string]interface{}
}

func (s RawJobLimitsImpl) JobLimits() BaseJobLimitsImpl {
	return s.jobLimits
}

func UnmarshalJobLimitsImplementation(input []byte) (JobLimits, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling JobLimits into map[string]interface: %+v", err)
	}

	value, ok := temp["jobLimitsType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Command") {
		var out CommandJobLimits
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CommandJobLimits: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Sweep") {
		var out SweepJobLimits
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SweepJobLimits: %+v", err)
		}
		return out, nil
	}

	var parent BaseJobLimitsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseJobLimitsImpl: %+v", err)
	}

	return RawJobLimitsImpl{
		jobLimits: parent,
		Type:      value,
		Values:    temp,
	}, nil

}
