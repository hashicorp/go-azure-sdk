package replicationevents

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventProperties struct {
	AffectedObjectCorrelationId *string                       `json:"affectedObjectCorrelationId,omitempty"`
	AffectedObjectFriendlyName  *string                       `json:"affectedObjectFriendlyName,omitempty"`
	Description                 *string                       `json:"description,omitempty"`
	EventCode                   *string                       `json:"eventCode,omitempty"`
	EventSpecificDetails        *EventSpecificDetails         `json:"eventSpecificDetails,omitempty"`
	EventType                   *string                       `json:"eventType,omitempty"`
	FabricId                    *string                       `json:"fabricId,omitempty"`
	HealthErrors                *[]HealthError                `json:"healthErrors,omitempty"`
	ProviderSpecificDetails     *EventProviderSpecificDetails `json:"providerSpecificDetails,omitempty"`
	Severity                    *string                       `json:"severity,omitempty"`
	TimeOfOccurrence            *string                       `json:"timeOfOccurrence,omitempty"`
}

func (o *EventProperties) GetTimeOfOccurrenceAsTime() (*time.Time, error) {
	if o.TimeOfOccurrence == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.TimeOfOccurrence, "2006-01-02T15:04:05Z07:00")
}

func (o *EventProperties) SetTimeOfOccurrenceAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.TimeOfOccurrence = &formatted
}

var _ json.Unmarshaler = &EventProperties{}

func (s *EventProperties) UnmarshalJSON(bytes []byte) error {
	type alias EventProperties
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into EventProperties: %+v", err)
	}

	s.AffectedObjectCorrelationId = decoded.AffectedObjectCorrelationId
	s.AffectedObjectFriendlyName = decoded.AffectedObjectFriendlyName
	s.Description = decoded.Description
	s.EventCode = decoded.EventCode
	s.EventType = decoded.EventType
	s.FabricId = decoded.FabricId
	s.HealthErrors = decoded.HealthErrors
	s.Severity = decoded.Severity
	s.TimeOfOccurrence = decoded.TimeOfOccurrence

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EventProperties into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["eventSpecificDetails"]; ok {
		impl, err := unmarshalEventSpecificDetailsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'EventSpecificDetails' for 'EventProperties': %+v", err)
		}
		s.EventSpecificDetails = impl
	}

	if v, ok := temp["providerSpecificDetails"]; ok {
		impl, err := unmarshalEventProviderSpecificDetailsImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'ProviderSpecificDetails' for 'EventProperties': %+v", err)
		}
		s.ProviderSpecificDetails = impl
	}
	return nil
}
