package reservationrecommendations

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ReservationRecommendation = ModernReservationRecommendation{}

type ModernReservationRecommendation struct {
	Properties ModernReservationRecommendationProperties `json:"properties"`

	// Fields inherited from ReservationRecommendation

	Etag     *string                       `json:"etag,omitempty"`
	Id       *string                       `json:"id,omitempty"`
	Kind     ReservationRecommendationKind `json:"kind"`
	Location *string                       `json:"location,omitempty"`
	Name     *string                       `json:"name,omitempty"`
	Sku      *string                       `json:"sku,omitempty"`
	Tags     *map[string]string            `json:"tags,omitempty"`
	Type     *string                       `json:"type,omitempty"`
}

func (s ModernReservationRecommendation) ReservationRecommendation() BaseReservationRecommendationImpl {
	return BaseReservationRecommendationImpl{
		Etag:     s.Etag,
		Id:       s.Id,
		Kind:     s.Kind,
		Location: s.Location,
		Name:     s.Name,
		Sku:      s.Sku,
		Tags:     s.Tags,
		Type:     s.Type,
	}
}

var _ json.Marshaler = ModernReservationRecommendation{}

func (s ModernReservationRecommendation) MarshalJSON() ([]byte, error) {
	type wrapper ModernReservationRecommendation
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ModernReservationRecommendation: %+v", err)
	}

	var decoded map[string]interface{}
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ModernReservationRecommendation: %+v", err)
	}

	decoded["kind"] = "modern"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ModernReservationRecommendation: %+v", err)
	}

	return encoded, nil
}

var _ json.Unmarshaler = &ModernReservationRecommendation{}

func (s *ModernReservationRecommendation) UnmarshalJSON(bytes []byte) error {
	var decoded struct {
		Etag     *string                       `json:"etag,omitempty"`
		Id       *string                       `json:"id,omitempty"`
		Kind     ReservationRecommendationKind `json:"kind"`
		Location *string                       `json:"location,omitempty"`
		Name     *string                       `json:"name,omitempty"`
		Sku      *string                       `json:"sku,omitempty"`
		Tags     *map[string]string            `json:"tags,omitempty"`
		Type     *string                       `json:"type,omitempty"`
	}
	if err := json.Unmarshal(bytes, &decoded); err != nil {
		return fmt.Errorf("unmarshaling: %+v", err)
	}

	s.Etag = decoded.Etag
	s.Id = decoded.Id
	s.Kind = decoded.Kind
	s.Location = decoded.Location
	s.Name = decoded.Name
	s.Sku = decoded.Sku
	s.Tags = decoded.Tags
	s.Type = decoded.Type

	var temp map[string]json.RawMessage
	if err := json.Unmarshal(bytes, &temp); err != nil {
		return fmt.Errorf("unmarshaling ModernReservationRecommendation into map[string]json.RawMessage: %+v", err)
	}

	if v, ok := temp["properties"]; ok {
		impl, err := UnmarshalModernReservationRecommendationPropertiesImplementation(v)
		if err != nil {
			return fmt.Errorf("unmarshaling field 'Properties' for 'ModernReservationRecommendation': %+v", err)
		}
		s.Properties = impl
	}

	return nil
}
