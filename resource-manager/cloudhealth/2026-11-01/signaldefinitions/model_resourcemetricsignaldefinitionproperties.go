package signaldefinitions

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SignalDefinitionProperties = ResourceMetricSignalDefinitionProperties{}

type ResourceMetricSignalDefinitionProperties struct {
	AggregationType MetricAggregationType `json:"aggregationType"`
	DimensionFilter *string               `json:"dimensionFilter,omitempty"`
	MetricName      string                `json:"metricName"`
	MetricNamespace string                `json:"metricNamespace"`
	TimeGrain       string                `json:"timeGrain"`

	// Fields inherited from SignalDefinitionProperties

	DataUnit          *string                       `json:"dataUnit,omitempty"`
	DisplayName       *string                       `json:"displayName,omitempty"`
	EvaluationRules   EvaluationRule                `json:"evaluationRules"`
	ProvisioningState *HealthModelProvisioningState `json:"provisioningState,omitempty"`
	RefreshInterval   *RefreshInterval              `json:"refreshInterval,omitempty"`
	SignalKind        SignalKind                    `json:"signalKind"`
	Tags              *map[string]string            `json:"tags,omitempty"`
}

func (s ResourceMetricSignalDefinitionProperties) SignalDefinitionProperties() BaseSignalDefinitionPropertiesImpl {
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

var _ json.Marshaler = ResourceMetricSignalDefinitionProperties{}

func (s ResourceMetricSignalDefinitionProperties) MarshalJSON() ([]byte, error) {
	type wrapper ResourceMetricSignalDefinitionProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ResourceMetricSignalDefinitionProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ResourceMetricSignalDefinitionProperties: %+v", err)
	}

	decoded["signalKind"] = "AzureResourceMetric"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ResourceMetricSignalDefinitionProperties: %+v", err)
	}

	return encoded, nil
}
