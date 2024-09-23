package alertrulesapis

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RuleDataSource interface {
	RuleDataSource() BaseRuleDataSourceImpl
}

var _ RuleDataSource = BaseRuleDataSourceImpl{}

type BaseRuleDataSourceImpl struct {
	LegacyResourceId *string `json:"legacyResourceId,omitempty"`
	MetricNamespace  *string `json:"metricNamespace,omitempty"`
	OdataType        string  `json:"odata.type"`
	ResourceLocation *string `json:"resourceLocation,omitempty"`
	ResourceUri      *string `json:"resourceUri,omitempty"`
}

func (s BaseRuleDataSourceImpl) RuleDataSource() BaseRuleDataSourceImpl {
	return s
}

var _ RuleDataSource = RawRuleDataSourceImpl{}

// RawRuleDataSourceImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawRuleDataSourceImpl struct {
	ruleDataSource BaseRuleDataSourceImpl
	Type           string
	Values         map[string]interface{}
}

func (s RawRuleDataSourceImpl) RuleDataSource() BaseRuleDataSourceImpl {
	return s.ruleDataSource
}

func UnmarshalRuleDataSourceImplementation(input []byte) (RuleDataSource, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling RuleDataSource into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["odata.type"]; ok {
		value = fmt.Sprintf("%v", v)
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

	var parent BaseRuleDataSourceImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseRuleDataSourceImpl: %+v", err)
	}

	return RawRuleDataSourceImpl{
		ruleDataSource: parent,
		Type:           value,
		Values:         temp,
	}, nil

}
