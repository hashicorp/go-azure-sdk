package schedule

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type Nodes interface {
	Nodes() BaseNodesImpl
}

var _ Nodes = BaseNodesImpl{}

type BaseNodesImpl struct {
	NodesValueType NodesValueType `json:"nodesValueType"`
}

func (s BaseNodesImpl) Nodes() BaseNodesImpl {
	return s
}

var _ Nodes = RawNodesImpl{}

// RawNodesImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawNodesImpl struct {
	nodes  BaseNodesImpl
	Type   string
	Values map[string]interface{}
}

func (s RawNodesImpl) Nodes() BaseNodesImpl {
	return s.nodes
}

func UnmarshalNodesImplementation(input []byte) (Nodes, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Nodes into map[string]interface: %+v", err)
	}

	value, ok := temp["nodesValueType"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "All") {
		var out AllNodes
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into AllNodes: %+v", err)
		}
		return out, nil
	}

	var parent BaseNodesImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseNodesImpl: %+v", err)
	}

	return RawNodesImpl{
		nodes:  parent,
		Type:   value,
		Values: temp,
	}, nil

}
