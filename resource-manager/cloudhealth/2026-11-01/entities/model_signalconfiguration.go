package entities

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SignalConfiguration struct {
	ApplicableResourceTypes *[]string                         `json:"applicableResourceTypes,omitempty"`
	Configuration           SignalRecommendationConfiguration `json:"configuration"`
	DataUnit                *string                           `json:"dataUnit,omitempty"`
	Description             *string                           `json:"description,omitempty"`
	DisplayName             *string                           `json:"displayName,omitempty"`
	EvaluationRules         *EvaluationRule                   `json:"evaluationRules,omitempty"`
	RefreshInterval         *RefreshInterval                  `json:"refreshInterval,omitempty"`
	SignalId                string                            `json:"signalId"`
}

var _ json.Unmarshaler = &SignalConfiguration{}

func (s *SignalConfiguration) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		ApplicableResourceTypes *[]string        `json:"applicableResourceTypes,omitempty"`
		DataUnit                *string          `json:"dataUnit,omitempty"`
		Description             *string          `json:"description,omitempty"`
		DisplayName             *string          `json:"displayName,omitempty"`
		EvaluationRules         *EvaluationRule  `json:"evaluationRules,omitempty"`
		RefreshInterval         *RefreshInterval `json:"refreshInterval,omitempty"`
		SignalId                string           `json:"signalId"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.ApplicableResourceTypes = decoded.ApplicableResourceTypes
	s.DataUnit = decoded.DataUnit
	s.Description = decoded.Description
	s.DisplayName = decoded.DisplayName
	s.EvaluationRules = decoded.EvaluationRules
	s.RefreshInterval = decoded.RefreshInterval
	s.SignalId = decoded.SignalId

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling SignalConfiguration into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["configuration"]; ok {
		impl, err := UnmarshalSignalRecommendationConfigurationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Configuration' for 'SignalConfiguration': %+v", err)
		}
		s.Configuration = impl
	}

	return nil
}
