package dataflow

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataflowDestinationHeaderAction = DataflowDestinationRemoveHeaderAction{}

type DataflowDestinationRemoveHeaderAction struct {
	Key string `json:"key"`

	// Fields inherited from DataflowDestinationHeaderAction

	ActionType DataflowHeaderActionType `json:"actionType"`
}

func (s DataflowDestinationRemoveHeaderAction) DataflowDestinationHeaderAction() BaseDataflowDestinationHeaderActionImpl {
	return BaseDataflowDestinationHeaderActionImpl{
		ActionType: s.ActionType,
	}
}

var _ json.Marshaler = DataflowDestinationRemoveHeaderAction{}

func (s DataflowDestinationRemoveHeaderAction) MarshalJSON() ([]byte, error) {
	type wrapper DataflowDestinationRemoveHeaderAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DataflowDestinationRemoveHeaderAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DataflowDestinationRemoveHeaderAction: %+v", err)
	}

	decoded["actionType"] = "Remove"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DataflowDestinationRemoveHeaderAction: %+v", err)
	}

	return encoded, nil
}
