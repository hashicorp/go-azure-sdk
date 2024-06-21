package reservationrecommendations

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ModernReservationRecommendationProperties = ModernSharedScopeReservationRecommendationProperties{}

type ModernSharedScopeReservationRecommendationProperties struct {

	// Fields inherited from ModernReservationRecommendationProperties
	CostWithNoReservedInstances    *Amount        `json:"costWithNoReservedInstances,omitempty"`
	FirstUsageDate                 *string        `json:"firstUsageDate,omitempty"`
	InstanceFlexibilityGroup       *string        `json:"instanceFlexibilityGroup,omitempty"`
	InstanceFlexibilityRatio       *float64       `json:"instanceFlexibilityRatio,omitempty"`
	LastUsageDate                  *string        `json:"lastUsageDate,omitempty"`
	Location                       *string        `json:"location,omitempty"`
	LookBackPeriod                 *int64         `json:"lookBackPeriod,omitempty"`
	MeterId                        *string        `json:"meterId,omitempty"`
	NetSavings                     *Amount        `json:"netSavings,omitempty"`
	NormalizedSize                 *string        `json:"normalizedSize,omitempty"`
	RecommendedQuantity            *float64       `json:"recommendedQuantity,omitempty"`
	RecommendedQuantityNormalized  *float64       `json:"recommendedQuantityNormalized,omitempty"`
	ResourceType                   *string        `json:"resourceType,omitempty"`
	SkuName                        *string        `json:"skuName,omitempty"`
	SkuProperties                  *[]SkuProperty `json:"skuProperties,omitempty"`
	Term                           *string        `json:"term,omitempty"`
	TotalCostWithReservedInstances *Amount        `json:"totalCostWithReservedInstances,omitempty"`
	TotalHours                     *int64         `json:"totalHours,omitempty"`
}

var _ json.Marshaler = ModernSharedScopeReservationRecommendationProperties{}

func (s ModernSharedScopeReservationRecommendationProperties) MarshalJSON() ([]byte, error) {
	type wrapper ModernSharedScopeReservationRecommendationProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ModernSharedScopeReservationRecommendationProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ModernSharedScopeReservationRecommendationProperties: %+v", err)
	}
	decoded["scope"] = "Shared"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ModernSharedScopeReservationRecommendationProperties: %+v", err)
	}

	return encoded, nil
}
