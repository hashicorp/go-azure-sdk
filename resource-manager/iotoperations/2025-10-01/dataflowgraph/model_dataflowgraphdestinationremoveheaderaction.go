package dataflowgraph

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataflowGraphDestinationHeaderAction = DataflowGraphDestinationRemoveHeaderAction{}

type DataflowGraphDestinationRemoveHeaderAction struct {
	Key string `json:"key"`

	// Fields inherited from DataflowGraphDestinationHeaderAction

	ActionType DataflowGraphDestinationHeaderActionType `json:"actionType"`
}

func (s DataflowGraphDestinationRemoveHeaderAction) DataflowGraphDestinationHeaderAction() BaseDataflowGraphDestinationHeaderActionImpl {
	return BaseDataflowGraphDestinationHeaderActionImpl{
		ActionType: s.ActionType,
	}
}

var _ json.Marshaler = DataflowGraphDestinationRemoveHeaderAction{}

func (s DataflowGraphDestinationRemoveHeaderAction) MarshalJSON() ([]byte, error) {
	type wrapper DataflowGraphDestinationRemoveHeaderAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DataflowGraphDestinationRemoveHeaderAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DataflowGraphDestinationRemoveHeaderAction: %+v", err)
	}

	decoded["actionType"] = "Remove"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DataflowGraphDestinationRemoveHeaderAction: %+v", err)
	}

	return encoded, nil
}
