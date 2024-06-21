package eventsubscriptions

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventSubscriptionUpdateParameters struct {
	DeadLetterDestination          DeadLetterDestination           `json:"deadLetterDestination"`
	DeadLetterWithResourceIdentity *DeadLetterWithResourceIdentity `json:"deadLetterWithResourceIdentity,omitempty"`
	DeliveryWithResourceIdentity   *DeliveryWithResourceIdentity   `json:"deliveryWithResourceIdentity,omitempty"`
	Destination                    EventSubscriptionDestination    `json:"destination"`
	EventDeliverySchema            *EventDeliverySchema            `json:"eventDeliverySchema,omitempty"`
	ExpirationTimeUtc              *string                         `json:"expirationTimeUtc,omitempty"`
	Filter                         *EventSubscriptionFilter        `json:"filter,omitempty"`
	Labels                         *[]string                       `json:"labels,omitempty"`
	RetryPolicy                    *RetryPolicy                    `json:"retryPolicy,omitempty"`
}

var _ json.Unmarshaler = &EventSubscriptionUpdateParameters{}

func (s *EventSubscriptionUpdateParameters) UnmarshalJSON(bytes []byte) error {
	type alias EventSubscriptionUpdateParameters
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into EventSubscriptionUpdateParameters: %+v", err)
	}

	s.DeadLetterWithResourceIdentity = decoded.DeadLetterWithResourceIdentity
	s.DeliveryWithResourceIdentity = decoded.DeliveryWithResourceIdentity
	s.EventDeliverySchema = decoded.EventDeliverySchema
	s.ExpirationTimeUtc = decoded.ExpirationTimeUtc
	s.Filter = decoded.Filter
	s.Labels = decoded.Labels
	s.RetryPolicy = decoded.RetryPolicy

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling EventSubscriptionUpdateParameters into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["deadLetterDestination"]; ok {
		impl, err := unmarshalDeadLetterDestinationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'DeadLetterDestination' for 'EventSubscriptionUpdateParameters': %+v", err)
		}
		s.DeadLetterDestination = impl
	}

	if v, ok := temp["destination"]; ok {
		impl, err := unmarshalEventSubscriptionDestinationImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Destination' for 'EventSubscriptionUpdateParameters': %+v", err)
		}
		s.Destination = impl
	}
	return nil
}
