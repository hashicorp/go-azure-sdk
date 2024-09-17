package reservationrecommendations

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
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
	Scope                          string         `json:"scope"`
	SkuName                        *string        `json:"skuName,omitempty"`
	SkuProperties                  *[]SkuProperty `json:"skuProperties,omitempty"`
	Term                           *string        `json:"term,omitempty"`
	TotalCostWithReservedInstances *Amount        `json:"totalCostWithReservedInstances,omitempty"`
	TotalHours                     *int64         `json:"totalHours,omitempty"`
}

func (s ModernSharedScopeReservationRecommendationProperties) ModernReservationRecommendationProperties() BaseModernReservationRecommendationPropertiesImpl {
	return BaseModernReservationRecommendationPropertiesImpl{
		CostWithNoReservedInstances:    s.CostWithNoReservedInstances,
		FirstUsageDate:                 s.FirstUsageDate,
		InstanceFlexibilityGroup:       s.InstanceFlexibilityGroup,
		InstanceFlexibilityRatio:       s.InstanceFlexibilityRatio,
		LastUsageDate:                  s.LastUsageDate,
		Location:                       s.Location,
		LookBackPeriod:                 s.LookBackPeriod,
		MeterId:                        s.MeterId,
		NetSavings:                     s.NetSavings,
		NormalizedSize:                 s.NormalizedSize,
		RecommendedQuantity:            s.RecommendedQuantity,
		RecommendedQuantityNormalized:  s.RecommendedQuantityNormalized,
		ResourceType:                   s.ResourceType,
		Scope:                          s.Scope,
		SkuName:                        s.SkuName,
		SkuProperties:                  s.SkuProperties,
		Term:                           s.Term,
		TotalCostWithReservedInstances: s.TotalCostWithReservedInstances,
		TotalHours:                     s.TotalHours,
	}
}

func (o *ModernSharedScopeReservationRecommendationProperties) GetFirstUsageDateAsTime() (*time.Time, error) {
	if o.FirstUsageDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.FirstUsageDate, "2006-01-02T15:04:05Z07:00")
}

func (o *ModernSharedScopeReservationRecommendationProperties) SetFirstUsageDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.FirstUsageDate = &formatted
}

func (o *ModernSharedScopeReservationRecommendationProperties) GetLastUsageDateAsTime() (*time.Time, error) {
	if o.LastUsageDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastUsageDate, "2006-01-02T15:04:05Z07:00")
}

func (o *ModernSharedScopeReservationRecommendationProperties) SetLastUsageDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.LastUsageDate = &formatted
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
	if err = json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling ModernSharedScopeReservationRecommendationProperties: %+v", err)
	}

	decoded["scope"] = "Shared"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling ModernSharedScopeReservationRecommendationProperties: %+v", err)
	}

	return encoded, nil
}
