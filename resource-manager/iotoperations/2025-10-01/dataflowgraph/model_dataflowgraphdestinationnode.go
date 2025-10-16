package dataflowgraph

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataflowGraphNode = DataflowGraphDestinationNode{}

type DataflowGraphDestinationNode struct {
	DestinationSettings DataflowGraphDestinationNodeSettings `json:"destinationSettings"`

	// Fields inherited from DataflowGraphNode

	Name     string                `json:"name"`
	NodeType DataflowGraphNodeType `json:"nodeType"`
}

func (s DataflowGraphDestinationNode) DataflowGraphNode() BaseDataflowGraphNodeImpl {
	return BaseDataflowGraphNodeImpl{
		Name:     s.Name,
		NodeType: s.NodeType,
	}
}

var _ json.Marshaler = DataflowGraphDestinationNode{}

func (s DataflowGraphDestinationNode) MarshalJSON() ([]byte, error) {
	type wrapper DataflowGraphDestinationNode
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DataflowGraphDestinationNode: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DataflowGraphDestinationNode: %+v", err)
	}

	decoded["nodeType"] = "Destination"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DataflowGraphDestinationNode: %+v", err)
	}

	return encoded, nil
}
