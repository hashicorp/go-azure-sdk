package entities

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SignalInstanceProperties = LogAnalyticsSignal{}

type LogAnalyticsSignal struct {
	DataUnit        *string          `json:"dataUnit,omitempty"`
	DisplayName     *string          `json:"displayName,omitempty"`
	EvaluationRules *EvaluationRule  `json:"evaluationRules,omitempty"`
	QueryText       *string          `json:"queryText,omitempty"`
	RefreshInterval *RefreshInterval `json:"refreshInterval,omitempty"`
	TimeGrain       *string          `json:"timeGrain,omitempty"`
	ValueColumnName *string          `json:"valueColumnName,omitempty"`

	// Fields inherited from SignalInstanceProperties

	Name                 string        `json:"name"`
	SignalDefinitionName *string       `json:"signalDefinitionName,omitempty"`
	SignalKind           SignalKind    `json:"signalKind"`
	Status               *SignalStatus `json:"status,omitempty"`
}

func (s LogAnalyticsSignal) SignalInstanceProperties() BaseSignalInstancePropertiesImpl {
	return BaseSignalInstancePropertiesImpl{
		Name:                 s.Name,
		SignalDefinitionName: s.SignalDefinitionName,
		SignalKind:           s.SignalKind,
		Status:               s.Status,
	}
}

var _ json.Marshaler = LogAnalyticsSignal{}

func (s LogAnalyticsSignal) MarshalJSON() ([]byte, error) {
	type wrapper LogAnalyticsSignal
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling LogAnalyticsSignal: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling LogAnalyticsSignal: %+v", err)
	}

	decoded["signalKind"] = "LogAnalyticsQuery"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling LogAnalyticsSignal: %+v", err)
	}

	return encoded, nil
}
