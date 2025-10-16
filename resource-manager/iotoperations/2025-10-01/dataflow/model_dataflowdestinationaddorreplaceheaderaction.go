package dataflow

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataflowDestinationHeaderAction = DataflowDestinationAddOrReplaceHeaderAction{}

type DataflowDestinationAddOrReplaceHeaderAction struct {
	Key   string `json:"key"`
	Value string `json:"value"`

	// Fields inherited from DataflowDestinationHeaderAction

	ActionType DataflowHeaderActionType `json:"actionType"`
}

func (s DataflowDestinationAddOrReplaceHeaderAction) DataflowDestinationHeaderAction() BaseDataflowDestinationHeaderActionImpl {
	return BaseDataflowDestinationHeaderActionImpl{
		ActionType: s.ActionType,
	}
}

var _ json.Marshaler = DataflowDestinationAddOrReplaceHeaderAction{}

func (s DataflowDestinationAddOrReplaceHeaderAction) MarshalJSON() ([]byte, error) {
	type wrapper DataflowDestinationAddOrReplaceHeaderAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DataflowDestinationAddOrReplaceHeaderAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DataflowDestinationAddOrReplaceHeaderAction: %+v", err)
	}

	decoded["actionType"] = "AddOrReplace"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DataflowDestinationAddOrReplaceHeaderAction: %+v", err)
	}

	return encoded, nil
}
