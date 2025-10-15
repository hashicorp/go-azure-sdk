package dataflowgraph

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataflowGraphDestinationHeaderAction interface {
	DataflowGraphDestinationHeaderAction() BaseDataflowGraphDestinationHeaderActionImpl
}

var _ DataflowGraphDestinationHeaderAction = BaseDataflowGraphDestinationHeaderActionImpl{}

type BaseDataflowGraphDestinationHeaderActionImpl struct {
	ActionType DataflowGraphDestinationHeaderActionType `json:"actionType"`
}

func (s BaseDataflowGraphDestinationHeaderActionImpl) DataflowGraphDestinationHeaderAction() BaseDataflowGraphDestinationHeaderActionImpl {
	return s
}

var _ DataflowGraphDestinationHeaderAction = RawDataflowGraphDestinationHeaderActionImpl{}

// RawDataflowGraphDestinationHeaderActionImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDataflowGraphDestinationHeaderActionImpl struct {
	dataflowGraphDestinationHeaderAction BaseDataflowGraphDestinationHeaderActionImpl
	Type                                 string
	Values                               map[string]interface{}
}

func (s RawDataflowGraphDestinationHeaderActionImpl) DataflowGraphDestinationHeaderAction() BaseDataflowGraphDestinationHeaderActionImpl {
	return s.dataflowGraphDestinationHeaderAction
}

func UnmarshalDataflowGraphDestinationHeaderActionImplementation(input []byte) (DataflowGraphDestinationHeaderAction, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DataflowGraphDestinationHeaderAction into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["actionType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "AddIfNotPresent") {
		var out DataflowGraphDestinationAddIfNotPresentHeaderAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DataflowGraphDestinationAddIfNotPresentHeaderAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "AddOrReplace") {
		var out DataflowGraphDestinationAddOrReplaceHeaderAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DataflowGraphDestinationAddOrReplaceHeaderAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Remove") {
		var out DataflowGraphDestinationRemoveHeaderAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DataflowGraphDestinationRemoveHeaderAction: %+v", err)
		}
		return out, nil
	}

	var parent BaseDataflowGraphDestinationHeaderActionImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDataflowGraphDestinationHeaderActionImpl: %+v", err)
	}

	return RawDataflowGraphDestinationHeaderActionImpl{
		dataflowGraphDestinationHeaderAction: parent,
		Type:                                 value,
		Values:                               temp,
	}, nil

}
