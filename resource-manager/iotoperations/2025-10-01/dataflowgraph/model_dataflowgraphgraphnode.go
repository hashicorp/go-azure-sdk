package dataflowgraph

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataflowGraphNode = DataflowGraphGraphNode{}

type DataflowGraphGraphNode struct {
	GraphSettings DataflowGraphNodeGraphSettings `json:"graphSettings"`

	// Fields inherited from DataflowGraphNode

	Name     string                `json:"name"`
	NodeType DataflowGraphNodeType `json:"nodeType"`
}

func (s DataflowGraphGraphNode) DataflowGraphNode() BaseDataflowGraphNodeImpl {
	return BaseDataflowGraphNodeImpl{
		Name:     s.Name,
		NodeType: s.NodeType,
	}
}

var _ json.Marshaler = DataflowGraphGraphNode{}

func (s DataflowGraphGraphNode) MarshalJSON() ([]byte, error) {
	type wrapper DataflowGraphGraphNode
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DataflowGraphGraphNode: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DataflowGraphGraphNode: %+v", err)
	}

	decoded["nodeType"] = "Graph"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DataflowGraphGraphNode: %+v", err)
	}

	return encoded, nil
}
