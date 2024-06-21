package reservationrecommendations

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ LegacyReservationRecommendationProperties = LegacySharedScopeReservationRecommendationProperties{}

type LegacySharedScopeReservationRecommendationProperties struct {

	// Fields inherited from LegacyReservationRecommendationProperties
	CostWithNoReservedInstances    *float64       `json:"costWithNoReservedInstances,omitempty"`
	FirstUsageDate                 *string        `json:"firstUsageDate,omitempty"`
	InstanceFlexibilityGroup       *string        `json:"instanceFlexibilityGroup,omitempty"`
	InstanceFlexibilityRatio       *float64       `json:"instanceFlexibilityRatio,omitempty"`
	LookBackPeriod                 *string        `json:"lookBackPeriod,omitempty"`
	MeterId                        *string        `json:"meterId,omitempty"`
	NetSavings                     *float64       `json:"netSavings,omitempty"`
	NormalizedSize                 *string        `json:"normalizedSize,omitempty"`
	RecommendedQuantity            *float64       `json:"recommendedQuantity,omitempty"`
	RecommendedQuantityNormalized  *float64       `json:"recommendedQuantityNormalized,omitempty"`
	ResourceType                   *string        `json:"resourceType,omitempty"`
	SkuProperties                  *[]SkuProperty `json:"skuProperties,omitempty"`
	Term                           *string        `json:"term,omitempty"`
	TotalCostWithReservedInstances *float64       `json:"totalCostWithReservedInstances,omitempty"`
}

var _ json.Marshaler = LegacySharedScopeReservationRecommendationProperties{}

func (s LegacySharedScopeReservationRecommendationProperties) MarshalJSON() ([]byte, error) {
	type wrapper LegacySharedScopeReservationRecommendationProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling LegacySharedScopeReservationRecommendationProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling LegacySharedScopeReservationRecommendationProperties: %+v", err)
	}
	decoded["scope"] = "Shared"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling LegacySharedScopeReservationRecommendationProperties: %+v", err)
	}

	return encoded, nil
}
