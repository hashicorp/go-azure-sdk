package job

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

// RawNodesImpl is returned when the Discriminated Value doesn't match any of the defined types.
// It can also be used as a Request Payload to provide a raw JSON payload, which is useful
// for preserving arbitrary/extensible JSON properties across a round-trip.
type RawNodesImpl struct {
	nodes  BaseNodesImpl
	Type   string
	Values map[string]interface{}
}

func (s RawNodesImpl) Nodes() BaseNodesImpl {
	return s.nodes
}

func (s RawNodesImpl) MarshalJSON() ([]byte, error) {
	return json.Marshal(s.Values)
}

func UnmarshalNodesImplementation(input []byte) (Nodes, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling Nodes into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["nodesValueType"]; ok {
		value = fmt.Sprintf("%v", v)
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
