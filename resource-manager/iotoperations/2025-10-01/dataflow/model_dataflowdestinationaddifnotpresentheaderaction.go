package dataflow

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ DataflowDestinationHeaderAction = DataflowDestinationAddIfNotPresentHeaderAction{}

type DataflowDestinationAddIfNotPresentHeaderAction struct {
	Key   string `json:"key"`
	Value string `json:"value"`

	// Fields inherited from DataflowDestinationHeaderAction

	ActionType DataflowHeaderActionType `json:"actionType"`
}

func (s DataflowDestinationAddIfNotPresentHeaderAction) DataflowDestinationHeaderAction() BaseDataflowDestinationHeaderActionImpl {
	return BaseDataflowDestinationHeaderActionImpl{
		ActionType: s.ActionType,
	}
}

var _ json.Marshaler = DataflowDestinationAddIfNotPresentHeaderAction{}

func (s DataflowDestinationAddIfNotPresentHeaderAction) MarshalJSON() ([]byte, error) {
	type wrapper DataflowDestinationAddIfNotPresentHeaderAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling DataflowDestinationAddIfNotPresentHeaderAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling DataflowDestinationAddIfNotPresentHeaderAction: %+v", err)
	}

	decoded["actionType"] = "AddIfNotPresent"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling DataflowDestinationAddIfNotPresentHeaderAction: %+v", err)
	}

	return encoded, nil
}
