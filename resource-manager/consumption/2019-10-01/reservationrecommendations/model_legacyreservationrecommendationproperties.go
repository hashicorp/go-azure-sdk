package reservationrecommendations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LegacyReservationRecommendationProperties struct {
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
	Scope                          *string        `json:"scope,omitempty"`
	SkuProperties                  *[]SkuProperty `json:"skuProperties,omitempty"`
	Term                           *string        `json:"term,omitempty"`
	TotalCostWithReservedInstances *float64       `json:"totalCostWithReservedInstances,omitempty"`
}
