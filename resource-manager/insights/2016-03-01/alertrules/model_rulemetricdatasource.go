package alertrules

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RuleDataSource = RuleMetricDataSource{}

type RuleMetricDataSource struct {
	MetricName *string `json:"metricName,omitempty"`

	// Fields inherited from RuleDataSource
	LegacyResourceId *string `json:"legacyResourceId,omitempty"`
	MetricNamespace  *string `json:"metricNamespace,omitempty"`
	ResourceLocation *string `json:"resourceLocation,omitempty"`
	ResourceUri      *string `json:"resourceUri,omitempty"`
}

var _ json.Marshaler = RuleMetricDataSource{}

func (s RuleMetricDataSource) MarshalJSON() ([]byte, error) {
	type wrapper RuleMetricDataSource
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RuleMetricDataSource: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RuleMetricDataSource: %+v", err)
	}
	decoded["odata.type"] = "Microsoft.Azure.Management.Insights.Models.RuleMetricDataSource"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RuleMetricDataSource: %+v", err)
	}

	return encoded, nil
}
