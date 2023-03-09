package replicationjobs

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ TaskTypeDetails = AutomationRunbookTaskDetails{}

type AutomationRunbookTaskDetails struct {
	AccountName         *string `json:"accountName,omitempty"`
	CloudServiceName    *string `json:"cloudServiceName,omitempty"`
	IsPrimarySideScript *bool   `json:"isPrimarySideScript,omitempty"`
	JobId               *string `json:"jobId,omitempty"`
	JobOutput           *string `json:"jobOutput,omitempty"`
	Name                *string `json:"name,omitempty"`
	RunbookId           *string `json:"runbookId,omitempty"`
	RunbookName         *string `json:"runbookName,omitempty"`
	SubscriptionId      *string `json:"subscriptionId,omitempty"`

	// Fields inherited from TaskTypeDetails
}

var _ json.Marshaler = AutomationRunbookTaskDetails{}

func (s AutomationRunbookTaskDetails) MarshalJSON() ([]byte, error) {
	type wrapper AutomationRunbookTaskDetails
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AutomationRunbookTaskDetails: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AutomationRunbookTaskDetails: %+v", err)
	}
	decoded["instanceType"] = "AutomationRunbookTaskDetails"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AutomationRunbookTaskDetails: %+v", err)
	}

	return encoded, nil
}
