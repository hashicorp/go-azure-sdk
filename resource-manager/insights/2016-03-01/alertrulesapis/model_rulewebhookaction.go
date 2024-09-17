package alertrulesapis

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RuleAction = RuleWebhookAction{}

type RuleWebhookAction struct {
	Properties *map[string]string `json:"properties,omitempty"`
	ServiceUri *string            `json:"serviceUri,omitempty"`

	// Fields inherited from RuleAction

	OdataType string `json:"odata.type"`
}

func (s RuleWebhookAction) RuleAction() BaseRuleActionImpl {
	return BaseRuleActionImpl{
		OdataType: s.OdataType,
	}
}

var _ json.Marshaler = RuleWebhookAction{}

func (s RuleWebhookAction) MarshalJSON() ([]byte, error) {
	type wrapper RuleWebhookAction
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RuleWebhookAction: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RuleWebhookAction: %+v", err)
	}

	decoded["odata.type"] = "Microsoft.Azure.Management.Insights.Models.RuleWebhookAction"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RuleWebhookAction: %+v", err)
	}

	return encoded, nil
}
