package entities

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ SignalInstanceProperties = ExternalSignal{}

type ExternalSignal struct {
	EvaluationRules *EvaluationRule `json:"evaluationRules,omitempty"`

	// Fields inherited from SignalInstanceProperties

	Name                 string        `json:"name"`
	SignalDefinitionName *string       `json:"signalDefinitionName,omitempty"`
	SignalKind           SignalKind    `json:"signalKind"`
	Status               *SignalStatus `json:"status,omitempty"`
}

func (s ExternalSignal) SignalInstanceProperties() BaseSignalInstancePropertiesImpl {
	return BaseSignalInstancePropertiesImpl{
		Name:                 s.Name,
		SignalDefinitionName: s.SignalDefinitionName,
		SignalKind:           s.SignalKind,
		Status:               s.Status,
	}
}

var _ json.Marshaler = ExternalSignal{}

func (s ExternalSignal) MarshalJSON() ([]byte, error) {
	type wrapper ExternalSignal
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ExternalSignal: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ExternalSignal: %+v", err)
	}

	decoded["signalKind"] = "External"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ExternalSignal: %+v", err)
	}

	return encoded, nil
}
