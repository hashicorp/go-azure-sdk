package replicationjobs

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type GroupTaskDetails interface {
	GroupTaskDetails() BaseGroupTaskDetailsImpl
}

var _ GroupTaskDetails = BaseGroupTaskDetailsImpl{}

type BaseGroupTaskDetailsImpl struct {
	ChildTasks   *[]ASRTask `json:"childTasks,omitempty"`
	InstanceType string     `json:"instanceType"`
}

func (s BaseGroupTaskDetailsImpl) GroupTaskDetails() BaseGroupTaskDetailsImpl {
	return s
}

var _ GroupTaskDetails = RawGroupTaskDetailsImpl{}

// RawGroupTaskDetailsImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawGroupTaskDetailsImpl struct {
	groupTaskDetails BaseGroupTaskDetailsImpl
	Type             string
	Values           map[string]interface{}
}

func (s RawGroupTaskDetailsImpl) GroupTaskDetails() BaseGroupTaskDetailsImpl {
	return s.groupTaskDetails
}

func UnmarshalGroupTaskDetailsImplementation(input []byte) (GroupTaskDetails, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling GroupTaskDetails into map[string]interface: %+v", err)
	}

	value, ok := temp["instanceType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "InlineWorkflowTaskDetails") {
		var out InlineWorkflowTaskDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into InlineWorkflowTaskDetails: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "RecoveryPlanGroupTaskDetails") {
		var out RecoveryPlanGroupTaskDetails
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RecoveryPlanGroupTaskDetails: %+v", err)
		}
		return out, nil
	}

	var parent BaseGroupTaskDetailsImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseGroupTaskDetailsImpl: %+v", err)
	}

	return RawGroupTaskDetailsImpl{
		groupTaskDetails: parent,
		Type:             value,
		Values:           temp,
	}, nil

}
