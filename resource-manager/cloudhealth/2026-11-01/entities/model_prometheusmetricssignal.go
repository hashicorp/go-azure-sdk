package entities

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SignalInstanceProperties = PrometheusMetricsSignal{}

type PrometheusMetricsSignal struct {
	DataUnit        *string          `json:"dataUnit,omitempty"`
	DisplayName     *string          `json:"displayName,omitempty"`
	EvaluationRules *EvaluationRule  `json:"evaluationRules,omitempty"`
	QueryText       *string          `json:"queryText,omitempty"`
	RefreshInterval *RefreshInterval `json:"refreshInterval,omitempty"`
	TimeGrain       *string          `json:"timeGrain,omitempty"`

	// Fields inherited from SignalInstanceProperties

	Name                 string        `json:"name"`
	SignalDefinitionName *string       `json:"signalDefinitionName,omitempty"`
	SignalKind           SignalKind    `json:"signalKind"`
	Status               *SignalStatus `json:"status,omitempty"`
}

func (s PrometheusMetricsSignal) SignalInstanceProperties() BaseSignalInstancePropertiesImpl {
	return BaseSignalInstancePropertiesImpl{
		Name:                 s.Name,
		SignalDefinitionName: s.SignalDefinitionName,
		SignalKind:           s.SignalKind,
		Status:               s.Status,
	}
}

var _ json.Marshaler = PrometheusMetricsSignal{}

func (s PrometheusMetricsSignal) MarshalJSON() ([]byte, error) {
	type wrapper PrometheusMetricsSignal
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling PrometheusMetricsSignal: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling PrometheusMetricsSignal: %+v", err)
	}

	decoded["signalKind"] = "PrometheusMetricsQuery"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling PrometheusMetricsSignal: %+v", err)
	}

	return encoded, nil
}
