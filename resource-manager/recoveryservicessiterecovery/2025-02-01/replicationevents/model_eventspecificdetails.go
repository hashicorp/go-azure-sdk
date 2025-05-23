package replicationevents

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventSpecificDetails interface {
	EventSpecificDetails() BaseEventSpecificDetailsImpl
}

var _ EventSpecificDetails = BaseEventSpecificDetailsImpl{}

type BaseEventSpecificDetailsImpl struct {
	InstanceType string `json:"instanceType"`
}

func (s BaseEventSpecificDetailsImpl) EventSpecificDetails() BaseEventSpecificDetailsImpl {
	return s
}

var _ EventSpecificDetails = RawEventSpecificDetailsImpl{}

// RawEventSpecificDetailsImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawEventSpecificDetailsImpl struct {
	eventSpecificDetails BaseEventSpecificDetailsImpl
	Type                 string
	Values               map[string]interface{}
}

func (s RawEventSpecificDetailsImpl) EventSpecificDetails() BaseEventSpecificDetailsImpl {
	return s.eventSpecificDetails
}

func UnmarshalEventSpecificDetailsImplementation(input []byte) (EventSpecificDetails, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling EventSpecificDetails into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["instanceType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "JobStatus") {
		var out JobStatusEventDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into JobStatusEventDetails: %+v", err)
		}
		return out, nil
	}

	var parent BaseEventSpecificDetailsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseEventSpecificDetailsImpl: %+v", err)
	}

	return RawEventSpecificDetailsImpl{
		eventSpecificDetails: parent,
		Type:                 value,
		Values:               temp,
	}, nil

}
