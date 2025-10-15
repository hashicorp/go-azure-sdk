package dataflowgraph

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DataflowGraphNode interface {
	DataflowGraphNode() BaseDataflowGraphNodeImpl
}

var _ DataflowGraphNode = BaseDataflowGraphNodeImpl{}

type BaseDataflowGraphNodeImpl struct {
	Name     string                `json:"name"`
	NodeType DataflowGraphNodeType `json:"nodeType"`
}

func (s BaseDataflowGraphNodeImpl) DataflowGraphNode() BaseDataflowGraphNodeImpl {
	return s
}

var _ DataflowGraphNode = RawDataflowGraphNodeImpl{}

// RawDataflowGraphNodeImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawDataflowGraphNodeImpl struct {
	dataflowGraphNode BaseDataflowGraphNodeImpl
	Type              string
	Values            map[string]interface{}
}

func (s RawDataflowGraphNodeImpl) DataflowGraphNode() BaseDataflowGraphNodeImpl {
	return s.dataflowGraphNode
}

func UnmarshalDataflowGraphNodeImplementation(input []byte) (DataflowGraphNode, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling DataflowGraphNode into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["nodeType"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Destination") {
		var out DataflowGraphDestinationNode
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DataflowGraphDestinationNode: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Graph") {
		var out DataflowGraphGraphNode
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DataflowGraphGraphNode: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Source") {
		var out DataflowGraphSourceNode
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into DataflowGraphSourceNode: %+v", err)
		}
		return out, nil
	}

	var parent BaseDataflowGraphNodeImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseDataflowGraphNodeImpl: %+v", err)
	}

	return RawDataflowGraphNodeImpl{
		dataflowGraphNode: parent,
		Type:              value,
		Values:            temp,
	}, nil

}
