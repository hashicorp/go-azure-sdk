package reservationrecommendations

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationRecommendation interface {
	ReservationRecommendation() BaseReservationRecommendationImpl
}

var _ ReservationRecommendation = BaseReservationRecommendationImpl{}

type BaseReservationRecommendationImpl struct {
	ETag     *string                       `json:"eTag,omitempty"`
	Id       *string                       `json:"id,omitempty"`
	Kind     ReservationRecommendationKind `json:"kind"`
	Location *string                       `json:"location,omitempty"`
	Name     *string                       `json:"name,omitempty"`
	Sku      *string                       `json:"sku,omitempty"`
	Tags     *map[string]string            `json:"tags,omitempty"`
	Type     *string                       `json:"type,omitempty"`
}

func (s BaseReservationRecommendationImpl) ReservationRecommendation() BaseReservationRecommendationImpl {
	return s
}

var _ ReservationRecommendation = RawReservationRecommendationImpl{}

// RawReservationRecommendationImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawReservationRecommendationImpl struct {
	reservationRecommendation BaseReservationRecommendationImpl
	Type                      string
	Values                    map[string]interface{}
}

func (s RawReservationRecommendationImpl) ReservationRecommendation() BaseReservationRecommendationImpl {
	return s.reservationRecommendation
}

func UnmarshalReservationRecommendationImplementation(input []byte) (ReservationRecommendation, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ReservationRecommendation into map[string]interface: %+v", err)
	}

	value, ok := temp["kind"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "legacy") {
		var out LegacyReservationRecommendation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LegacyReservationRecommendation: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "modern") {
		var out ModernReservationRecommendation
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ModernReservationRecommendation: %+v", err)
		}
		return out, nil
	}

	var parent BaseReservationRecommendationImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseReservationRecommendationImpl: %+v", err)
	}

	return RawReservationRecommendationImpl{
		reservationRecommendation: parent,
		Type:                      value,
		Values:                    temp,
	}, nil

}
