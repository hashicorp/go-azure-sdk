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
	ETag     *string            `json:"eTag,omitempty"`
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
