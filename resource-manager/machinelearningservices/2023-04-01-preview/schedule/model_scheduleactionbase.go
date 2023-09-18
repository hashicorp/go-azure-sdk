package schedule

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ScheduleActionBase interface {
}

// RawScheduleActionBaseImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawScheduleActionBaseImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalScheduleActionBaseImplementation(input []byte) (ScheduleActionBase, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ScheduleActionBase into map[string]interface: %+v", err)
	}

	value, ok := temp["actionType"].(string)
	if !ok {
		return nil, nil
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

	if strings.EqualFold(value, "ImportData") {
		var out ImportDataAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ImportDataAction: %+v", err)
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

	out := RawScheduleActionBaseImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
