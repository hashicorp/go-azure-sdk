package dataflowgraph

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataflowGraphDestinationHeaderAction = DataflowGraphDestinationAddIfNotPresentHeaderAction{}

type DataflowGraphDestinationAddIfNotPresentHeaderAction struct {
	Key   string `json:"key"`
	Value string `json:"value"`

	// Fields inherited from DataflowGraphDestinationHeaderAction

	ActionType DataflowGraphDestinationHeaderActionType `json:"actionType"`
}

func (s DataflowGraphDestinationAddIfNotPresentHeaderAction) DataflowGraphDestinationHeaderAction() BaseDataflowGraphDestinationHeaderActionImpl {
	return BaseDataflowGraphDestinationHeaderActionImpl{
		ActionType: s.ActionType,
	}
}

var _ json.Marshaler = DataflowGraphDestinationAddIfNotPresentHeaderAction{}

func (s DataflowGraphDestinationAddIfNotPresentHeaderAction) MarshalJSON() ([]byte, error) {
	type wrapper DataflowGraphDestinationAddIfNotPresentHeaderAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DataflowGraphDestinationAddIfNotPresentHeaderAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DataflowGraphDestinationAddIfNotPresentHeaderAction: %+v", err)
	}

	decoded["actionType"] = "AddIfNotPresent"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DataflowGraphDestinationAddIfNotPresentHeaderAction: %+v", err)
	}

	return encoded, nil
}
