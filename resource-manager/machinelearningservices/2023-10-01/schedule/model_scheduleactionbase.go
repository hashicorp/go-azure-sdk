package schedule

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduleActionBase interface {
	ScheduleActionBase() BaseScheduleActionBaseImpl
}

var _ ScheduleActionBase = BaseScheduleActionBaseImpl{}

type BaseScheduleActionBaseImpl struct {
	ActionType ScheduleActionType `json:"actionType"`
}

func (s BaseScheduleActionBaseImpl) ScheduleActionBase() BaseScheduleActionBaseImpl {
	return s
}

var _ ScheduleActionBase = RawScheduleActionBaseImpl{}

// RawScheduleActionBaseImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawScheduleActionBaseImpl struct {
	scheduleActionBase BaseScheduleActionBaseImpl
	Type               string
	Values             map[string]interface{}
}

func (s RawScheduleActionBaseImpl) ScheduleActionBase() BaseScheduleActionBaseImpl {
	return s.scheduleActionBase
}

func UnmarshalScheduleActionBaseImplementation(input []byte) (ScheduleActionBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ScheduleActionBase into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["actionType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "CreateMonitor") {
		var out CreateMonitorAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into CreateMonitorAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "InvokeBatchEndpoint") {
		var out EndpointScheduleAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into EndpointScheduleAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "CreateJob") {
		var out JobScheduleAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into JobScheduleAction: %+v", err)
		}
		return out, nil
	}

	var parent BaseScheduleActionBaseImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseScheduleActionBaseImpl: %+v", err)
	}

	return RawScheduleActionBaseImpl{
		scheduleActionBase: parent,
		Type:               value,
		Values:             temp,
	}, nil

}
