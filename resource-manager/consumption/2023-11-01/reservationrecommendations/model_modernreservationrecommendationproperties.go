package reservationrecommendations

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ModernReservationRecommendationProperties interface {
}

// RawModeOfTransitImpl is returned when the Discriminated Value
// doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawModernReservationRecommendationPropertiesImpl struct {
	Type   string
	Values map[string]interface{}
}

func unmarshalModernReservationRecommendationPropertiesImplementation(input []byte) (ModernReservationRecommendationProperties, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling ModernReservationRecommendationProperties into map[string]interface: %+v", err)
	}

	value, ok := temp["scope"].(string)
	if !ok {
		return nil, nil
	}

	if strings.EqualFold(value, "Shared") {
		var out ModernSharedScopeReservationRecommendationProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ModernSharedScopeReservationRecommendationProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Single") {
		var out ModernSingleScopeReservationRecommendationProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into ModernSingleScopeReservationRecommendationProperties: %+v", err)
		}
		return out, nil
	}

	out := RawModernReservationRecommendationPropertiesImpl{
		Type:   value,
		Values: temp,
	}
	return out, nil

}
