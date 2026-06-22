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

// RawGroupTaskDetailsImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawGroupTaskDetailsImpl struct {
	groupTaskDetails BaseGroupTaskDetailsImpl
	Type             string
	Values           map[string]interface{}
}

func (s RawGroupTaskDetailsImpl) GroupTaskDetails() BaseGroupTaskDetailsImpl {
	return s.groupTaskDetails
}

func (s RawGroupTaskDetailsImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalGroupTaskDetailsImplementation(input []byte) (GroupTaskDetails, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling GroupTaskDetails into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["instanceType"]; ok {
		value = fmt.Sprintf("%v", v)
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
