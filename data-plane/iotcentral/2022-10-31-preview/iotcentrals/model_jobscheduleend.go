package iotcentrals

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type JobScheduleEnd interface {
	JobScheduleEnd() BaseJobScheduleEndImpl
}

var _ JobScheduleEnd = BaseJobScheduleEndImpl{}

type BaseJobScheduleEndImpl struct {
	Type string `json:"type"`
}

func (s BaseJobScheduleEndImpl) JobScheduleEnd() BaseJobScheduleEndImpl {
	return s
}

var _ JobScheduleEnd = RawJobScheduleEndImpl{}

// RawJobScheduleEndImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawJobScheduleEndImpl struct {
	jobScheduleEnd BaseJobScheduleEndImpl
	Type           string
	Values         map[string]interface{}
}

func (s RawJobScheduleEndImpl) JobScheduleEnd() BaseJobScheduleEndImpl {
	return s.jobScheduleEnd
}

func UnmarshalJobScheduleEndImplementation(input []byte) (JobScheduleEnd, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling JobScheduleEnd into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["type"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "date") {
		var out DateJobScheduleEnd
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DateJobScheduleEnd: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "occurrences") {
		var out OccurrencesJobScheduleEnd
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into OccurrencesJobScheduleEnd: %+v", err)
		}
		return out, nil
	}

	var parent BaseJobScheduleEndImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseJobScheduleEndImpl: %+v", err)
	}

	return RawJobScheduleEndImpl{
		jobScheduleEnd: parent,
		Type:           value,
		Values:         temp,
	}, nil

}
