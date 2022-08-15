package reservationrecommendations

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ReservationRecommendation interface {
}

func unmarshalReservationRecommendationImplementation(input []byte) (ReservationRecommendation, error) {
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

	type RawReservationRecommendationImpl struct {
		Type   string                 `json:"-"`
		Values map[string]interface{} `json:"-"`
	}
	out := RawReservationRecommendationImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
