package alertrules

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RuleDataSource interface {
}

// RawModeOfTransitImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawRuleDataSourceImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalRuleDataSourceImplementation(input []byte) (RuleDataSource, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling RuleDataSource into map[string]interface: %+v", err)
	}

	value, ok := temp["odata.type"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Microsoft.Azure.Management.Insights.Models.RuleManagementEventDataSource") {
		var out RuleManagementEventDataSource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RuleManagementEventDataSource: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Microsoft.Azure.Management.Insights.Models.RuleMetricDataSource") {
		var out RuleMetricDataSource
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into RuleMetricDataSource: %+v", err)
		}
		return out, nil
	}

	out := RawRuleDataSourceImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
