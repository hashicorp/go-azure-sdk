package dataflow

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataflowDestinationHeaderAction interface {
	DataflowDestinationHeaderAction() BaseDataflowDestinationHeaderActionImpl
}

var _ DataflowDestinationHeaderAction = BaseDataflowDestinationHeaderActionImpl{}

type BaseDataflowDestinationHeaderActionImpl struct {
	ActionType DataflowHeaderActionType `json:"actionType"`
}

func (s BaseDataflowDestinationHeaderActionImpl) DataflowDestinationHeaderAction() BaseDataflowDestinationHeaderActionImpl {
	return s
}

var _ DataflowDestinationHeaderAction = RawDataflowDestinationHeaderActionImpl{}

// RawDataflowDestinationHeaderActionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDataflowDestinationHeaderActionImpl struct {
	dataflowDestinationHeaderAction BaseDataflowDestinationHeaderActionImpl
	Type                            string
	Values                          map[string]interface{}
}

func (s RawDataflowDestinationHeaderActionImpl) DataflowDestinationHeaderAction() BaseDataflowDestinationHeaderActionImpl {
	return s.dataflowDestinationHeaderAction
}

func UnmarshalDataflowDestinationHeaderActionImplementation(input []byte) (DataflowDestinationHeaderAction, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DataflowDestinationHeaderAction into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["actionType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "AddIfNotPresent") {
		var out DataflowDestinationAddIfNotPresentHeaderAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DataflowDestinationAddIfNotPresentHeaderAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AddOrReplace") {
		var out DataflowDestinationAddOrReplaceHeaderAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DataflowDestinationAddOrReplaceHeaderAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Remove") {
		var out DataflowDestinationRemoveHeaderAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DataflowDestinationRemoveHeaderAction: %+v", err)
		}
		return out, nil
	}

	var parent BaseDataflowDestinationHeaderActionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDataflowDestinationHeaderActionImpl: %+v", err)
	}

	return RawDataflowDestinationHeaderActionImpl{
		dataflowDestinationHeaderAction: parent,
		Type:                            value,
		Values:                          temp,
	}, nil

}
