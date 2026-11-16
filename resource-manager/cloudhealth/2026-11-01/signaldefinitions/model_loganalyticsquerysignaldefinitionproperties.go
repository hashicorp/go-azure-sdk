package signaldefinitions

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SignalDefinitionProperties = LogAnalyticsQuerySignalDefinitionProperties{}

type LogAnalyticsQuerySignalDefinitionProperties struct {
	QueryText       string  `json:"queryText"`
	TimeGrain       *string `json:"timeGrain,omitempty"`
	ValueColumnName *string `json:"valueColumnName,omitempty"`

	// Fields inherited from SignalDefinitionProperties

	DataUnit          *string                       `json:"dataUnit,omitempty"`
	DisplayName       *string                       `json:"displayName,omitempty"`
	EvaluationRules   EvaluationRule                `json:"evaluationRules"`
	ProvisioningState *HealthModelProvisioningState `json:"provisioningState,omitempty"`
	RefreshInterval   *RefreshInterval              `json:"refreshInterval,omitempty"`
	SignalKind        SignalKind                    `json:"signalKind"`
	Tags              *map[string]string            `json:"tags,omitempty"`
}

func (s LogAnalyticsQuerySignalDefinitionProperties) SignalDefinitionProperties() BaseSignalDefinitionPropertiesImpl {
	return BaseSignalDefinitionPropertiesImpl{
		DataUnit:          s.DataUnit,
		DisplayName:       s.DisplayName,
		EvaluationRules:   s.EvaluationRules,
		ProvisioningState: s.ProvisioningState,
		RefreshInterval:   s.RefreshInterval,
		SignalKind:        s.SignalKind,
		Tags:              s.Tags,
	}
}

var _ json.Marshaler = LogAnalyticsQuerySignalDefinitionProperties{}

func (s LogAnalyticsQuerySignalDefinitionProperties) MarshalJSON() ([]byte, error) {
	type wrapper LogAnalyticsQuerySignalDefinitionProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling LogAnalyticsQuerySignalDefinitionProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling LogAnalyticsQuerySignalDefinitionProperties: %+v", err)
	}

	decoded["signalKind"] = "LogAnalyticsQuery"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling LogAnalyticsQuerySignalDefinitionProperties: %+v", err)
	}

	return encoded, nil
}
