package dataflowgraph

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataflowGraphDestinationHeaderAction = DataflowGraphDestinationAddOrReplaceHeaderAction{}

type DataflowGraphDestinationAddOrReplaceHeaderAction struct {
	Key   string `json:"key"`
	Value string `json:"value"`

	// Fields inherited from DataflowGraphDestinationHeaderAction

	ActionType DataflowGraphDestinationHeaderActionType `json:"actionType"`
}

func (s DataflowGraphDestinationAddOrReplaceHeaderAction) DataflowGraphDestinationHeaderAction() BaseDataflowGraphDestinationHeaderActionImpl {
	return BaseDataflowGraphDestinationHeaderActionImpl{
		ActionType: s.ActionType,
	}
}

var _ json.Marshaler = DataflowGraphDestinationAddOrReplaceHeaderAction{}

func (s DataflowGraphDestinationAddOrReplaceHeaderAction) MarshalJSON() ([]byte, error) {
	type wrapper DataflowGraphDestinationAddOrReplaceHeaderAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DataflowGraphDestinationAddOrReplaceHeaderAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DataflowGraphDestinationAddOrReplaceHeaderAction: %+v", err)
	}

	decoded["actionType"] = "AddOrReplace"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DataflowGraphDestinationAddOrReplaceHeaderAction: %+v", err)
	}

	return encoded, nil
}
