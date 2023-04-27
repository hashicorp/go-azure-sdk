package reservationrecommendations

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ReservationRecommendation = LegacyReservationRecommendation{}

type LegacyReservationRecommendation struct {
	Properties LegacyReservationRecommendationProperties `json:"properties"`

	// Fields inherited from ReservationRecommendation
	Etag     *string            `json:"etag,omitempty"`
	Id       *string            `json:"id,omitempty"`
	Location *string            `json:"location,omitempty"`
	Name     *string            `json:"name,omitempty"`
	Sku      *string            `json:"sku,omitempty"`
	Tags     *map[string]string `json:"tags,omitempty"`
	Type     *string            `json:"type,omitempty"`
}

var _ json.Marshaler = LegacyReservationRecommendation{}

func (s LegacyReservationRecommendation) MarshalJSON() ([]byte, error) {
	type wrapper LegacyReservationRecommendation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling LegacyReservationRecommendation: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling LegacyReservationRecommendation: %+v", err)
	}
	decoded["kind"] = "legacy"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling LegacyReservationRecommendation: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &LegacyReservationRecommendation{}

func (s *LegacyReservationRecommendation) UnmarshalJSON(bytes []byte) error {
	type alias LegacyReservationRecommendation
	var decoded alias
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling into LegacyReservationRecommendation: %+v", err)
	}

	s.Etag = decoded.Etag
	s.Id = decoded.Id
	s.Location = decoded.Location
	s.Name = decoded.Name
	s.Sku = decoded.Sku
	s.Tags = decoded.Tags
	s.Type = decoded.Type

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling LegacyReservationRecommendation into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["properties"]; ok {
		impl, err := unmarshalLegacyReservationRecommendationPropertiesImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Properties' for 'LegacyReservationRecommendation': %+v", err)
		}
		s.Properties = impl
	}
	return nil
}
