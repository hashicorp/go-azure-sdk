package alertrules

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RuleAction = RuleEmailAction{}

type RuleEmailAction struct {
	CustomEmails        *[]string `json:"customEmails,omitempty"`
	SendToServiceOwners *bool     `json:"sendToServiceOwners,omitempty"`

	// Fields inherited from RuleAction
}

var _ json.Marshaler = RuleEmailAction{}

func (s RuleEmailAction) MarshalJSON() ([]byte, error) {
	type wrapper RuleEmailAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RuleEmailAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RuleEmailAction: %+v", err)
	}
	decoded["odata.type"] = "Microsoft.Azure.Management.Insights.Models.RuleEmailAction"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RuleEmailAction: %+v", err)
	}

	return encoded, nil
}
