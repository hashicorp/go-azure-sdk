package dataflowgraph

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataflowGraphNode = DataflowGraphSourceNode{}

type DataflowGraphSourceNode struct {
	SourceSettings DataflowGraphSourceSettings `json:"sourceSettings"`

	// Fields inherited from DataflowGraphNode

	Name     string                `json:"name"`
	NodeType DataflowGraphNodeType `json:"nodeType"`
}

func (s DataflowGraphSourceNode) DataflowGraphNode() BaseDataflowGraphNodeImpl {
	return BaseDataflowGraphNodeImpl{
		Name:     s.Name,
		NodeType: s.NodeType,
	}
}

var _ json.Marshaler = DataflowGraphSourceNode{}

func (s DataflowGraphSourceNode) MarshalJSON() ([]byte, error) {
	type wrapper DataflowGraphSourceNode
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DataflowGraphSourceNode: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DataflowGraphSourceNode: %+v", err)
	}

	decoded["nodeType"] = "Source"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DataflowGraphSourceNode: %+v", err)
	}

	return encoded, nil
}
