package alertrulesapis

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ RuleDataSource = RuleManagementEventDataSource{}

type RuleManagementEventDataSource struct {
	Claims               *RuleManagementEventClaimsDataSource `json:"claims,omitempty"`
	EventName            *string                              `json:"eventName,omitempty"`
	EventSource          *string                              `json:"eventSource,omitempty"`
	Level                *string                              `json:"level,omitempty"`
	OperationName        *string                              `json:"operationName,omitempty"`
	ResourceGroupName    *string                              `json:"resourceGroupName,omitempty"`
	ResourceProviderName *string                              `json:"resourceProviderName,omitempty"`
	Status               *string                              `json:"status,omitempty"`
	SubStatus            *string                              `json:"subStatus,omitempty"`

	// Fields inherited from RuleDataSource

	LegacyResourceId *string `json:"legacyResourceId,omitempty"`
	MetricNamespace  *string `json:"metricNamespace,omitempty"`
	OdataType        string  `json:"odata.type"`
	ResourceLocation *string `json:"resourceLocation,omitempty"`
	ResourceUri      *string `json:"resourceUri,omitempty"`
}

func (s RuleManagementEventDataSource) RuleDataSource() BaseRuleDataSourceImpl {
	return BaseRuleDataSourceImpl{
		LegacyResourceId: s.LegacyResourceId,
		MetricNamespace:  s.MetricNamespace,
		OdataType:        s.OdataType,
		ResourceLocation: s.ResourceLocation,
		ResourceUri:      s.ResourceUri,
	}
}

var _ json.Marshaler = RuleManagementEventDataSource{}

func (s RuleManagementEventDataSource) MarshalJSON() ([]byte, error) {
	type wrapper RuleManagementEventDataSource
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling RuleManagementEventDataSource: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling RuleManagementEventDataSource: %+v", err)
	}

	decoded["odata.type"] = "Microsoft.Azure.Management.Insights.Models.RuleManagementEventDataSource"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling RuleManagementEventDataSource: %+v", err)
	}

	return encoded, nil
}
