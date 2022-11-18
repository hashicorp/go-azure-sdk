package replicationjobs

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ GroupTaskDetails = InlineWorkflowTaskDetails{}

type InlineWorkflowTaskDetails struct {
	WorkflowIds *[]string `json:"workflowIds,omitempty"`

	// Fields inherited from GroupTaskDetails
	ChildTasks *[]ASRTask `json:"childTasks,omitempty"`
}

var _ json.Marshaler = InlineWorkflowTaskDetails{}

func (s InlineWorkflowTaskDetails) MarshalJSON() ([]byte, error) {
	type wrapper InlineWorkflowTaskDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling InlineWorkflowTaskDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling InlineWorkflowTaskDetails: %+v", err)
	}
	decoded["instanceType"] = "InlineWorkflowTaskDetails"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling InlineWorkflowTaskDetails: %+v", err)
	}

	return encoded, nil
}
