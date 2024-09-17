package reservationrecommendations

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ModernReservationRecommendationProperties interface {
	ModernReservationRecommendationProperties() BaseModernReservationRecommendationPropertiesImpl
}

var _ ModernReservationRecommendationProperties = BaseModernReservationRecommendationPropertiesImpl{}

type BaseModernReservationRecommendationPropertiesImpl struct {
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

func (s BaseModernReservationRecommendationPropertiesImpl) ModernReservationRecommendationProperties() BaseModernReservationRecommendationPropertiesImpl {
	return s
}

var _ ModernReservationRecommendationProperties = RawModernReservationRecommendationPropertiesImpl{}

// RawModernReservationRecommendationPropertiesImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawModernReservationRecommendationPropertiesImpl struct {
	modernReservationRecommendationProperties BaseModernReservationRecommendationPropertiesImpl
	Type                                      string
	Values                                    map[string]interface{}
}

func (s RawModernReservationRecommendationPropertiesImpl) ModernReservationRecommendationProperties() BaseModernReservationRecommendationPropertiesImpl {
	return s.modernReservationRecommendationProperties
}

func UnmarshalModernReservationRecommendationPropertiesImplementation(input []byte) (ModernReservationRecommendationProperties, error) {
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

	var parent BaseModernReservationRecommendationPropertiesImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseModernReservationRecommendationPropertiesImpl: %+v", err)
	}

	return RawModernReservationRecommendationPropertiesImpl{
		modernReservationRecommendationProperties: parent,
		Type:   value,
		Values: temp,
	}, nil

}
