package alertrules

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RuleAction interface {
}

// RawRuleActionImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawRuleActionImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalRuleActionImplementation(input []byte) (RuleAction, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling RuleAction into map[string]interface: %+v", err)
	}

	value, ok := temp["odata.type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Microsoft.Azure.Management.Insights.Models.RuleEmailAction") {
		var out RuleEmailAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RuleEmailAction: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Microsoft.Azure.Management.Insights.Models.RuleWebhookAction") {
		var out RuleWebhookAction
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RuleWebhookAction: %+v", err)
		}
		return out, nil
	}

	out := RawRuleActionImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
