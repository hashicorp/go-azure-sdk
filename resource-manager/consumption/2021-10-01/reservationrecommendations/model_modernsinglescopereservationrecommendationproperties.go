package reservationrecommendations

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ ModernReservationRecommendationProperties = ModernSingleScopeReservationRecommendationProperties{}

type ModernSingleScopeReservationRecommendationProperties struct {
	SubscriptionId *string `json:"subscriptionId,omitempty"`

	// Fields inherited from ModernReservationRecommendationProperties
	CostWithNoReservedInstances    *Amount        `json:"costWithNoReservedInstances,omitempty"`
	FirstUsageDate                 *string        `json:"firstUsageDate,omitempty"`
	InstanceFlexibilityGroup       *string        `json:"instanceFlexibilityGroup,omitempty"`
	InstanceFlexibilityRatio       *float64       `json:"instanceFlexibilityRatio,omitempty"`
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
}

func (o *ModernSingleScopeReservationRecommendationProperties) GetFirstUsageDateAsTime() (*time.Time, error) {
	if o.FirstUsageDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.FirstUsageDate, "2006-01-02T15:04:05Z07:00")
}

var _ json.Marshaler = ModernSingleScopeReservationRecommendationProperties{}

func (s ModernSingleScopeReservationRecommendationProperties) MarshalJSON() ([]byte, error) {
	type wrapper ModernSingleScopeReservationRecommendationProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling ModernSingleScopeReservationRecommendationProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ModernSingleScopeReservationRecommendationProperties: %+v", err)
	}
	decoded["scope"] = "Single"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ModernSingleScopeReservationRecommendationProperties: %+v", err)
	}

	return encoded, nil
}
