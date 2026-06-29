package replicationjobs

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TaskTypeDetails interface {
	TaskTypeDetails() BaseTaskTypeDetailsImpl
}

var _ TaskTypeDetails = BaseTaskTypeDetailsImpl{}

type BaseTaskTypeDetailsImpl struct {
	InstanceType string `json:"instanceType"`
}

func (s BaseTaskTypeDetailsImpl) TaskTypeDetails() BaseTaskTypeDetailsImpl {
	return s
}

var _ TaskTypeDetails = RawTaskTypeDetailsImpl{}

// RawTaskTypeDetailsImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawTaskTypeDetailsImpl struct {
	taskTypeDetails BaseTaskTypeDetailsImpl
	Type            string
	Values          map[string]interface{}
}

func (s RawTaskTypeDetailsImpl) TaskTypeDetails() BaseTaskTypeDetailsImpl {
	return s.taskTypeDetails
}

func (s RawTaskTypeDetailsImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalTaskTypeDetailsImplementation(input []byte) (TaskTypeDetails, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling TaskTypeDetails into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["instanceType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "AutomationRunbookTaskDetails") {
		var out AutomationRunbookTaskDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AutomationRunbookTaskDetails: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ConsistencyCheckTaskDetails") {
		var out ConsistencyCheckTaskDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ConsistencyCheckTaskDetails: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "JobTaskDetails") {
		var out JobTaskDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into JobTaskDetails: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ManualActionTaskDetails") {
		var out ManualActionTaskDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ManualActionTaskDetails: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "ScriptActionTaskDetails") {
		var out ScriptActionTaskDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ScriptActionTaskDetails: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "VmNicUpdatesTaskDetails") {
		var out VMNicUpdatesTaskDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into VMNicUpdatesTaskDetails: %+v", err)
		}
		return out, nil
	}

	var parent BaseTaskTypeDetailsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseTaskTypeDetailsImpl: %+v", err)
	}

	return RawTaskTypeDetailsImpl{
		taskTypeDetails: parent,
		Type:            value,
		Values:          temp,
	}, nil

}
