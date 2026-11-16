package entities

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SignalInstanceProperties = AzureResourceSignal{}

type AzureResourceSignal struct {
	AggregationType *MetricAggregationType `json:"aggregationType,omitempty"`
	DataUnit        *string                `json:"dataUnit,omitempty"`
	DimensionFilter *string                `json:"dimensionFilter,omitempty"`
	DisplayName     *string                `json:"displayName,omitempty"`
	EvaluationRules *EvaluationRule        `json:"evaluationRules,omitempty"`
	MetricName      *string                `json:"metricName,omitempty"`
	MetricNamespace *string                `json:"metricNamespace,omitempty"`
	RefreshInterval *RefreshInterval       `json:"refreshInterval,omitempty"`
	TimeGrain       *string                `json:"timeGrain,omitempty"`

	// Fields inherited from SignalInstanceProperties

	Name                 string        `json:"name"`
	SignalDefinitionName *string       `json:"signalDefinitionName,omitempty"`
	SignalKind           SignalKind    `json:"signalKind"`
	Status               *SignalStatus `json:"status,omitempty"`
}

func (s AzureResourceSignal) SignalInstanceProperties() BaseSignalInstancePropertiesImpl {
	return BaseSignalInstancePropertiesImpl{
		Name:                 s.Name,
		SignalDefinitionName: s.SignalDefinitionName,
		SignalKind:           s.SignalKind,
		Status:               s.Status,
	}
}

var _ json.Marshaler = AzureResourceSignal{}

func (s AzureResourceSignal) MarshalJSON() ([]byte, error) {
	type wrapper AzureResourceSignal
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling AzureResourceSignal: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling AzureResourceSignal: %+v", err)
	}

	decoded["signalKind"] = "AzureResourceMetric"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling AzureResourceSignal: %+v", err)
	}

	return encoded, nil
}
