package openapis

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LegacyReservationRecommendationProperties interface {
	LegacyReservationRecommendationProperties() BaseLegacyReservationRecommendationPropertiesImpl
}

var _ LegacyReservationRecommendationProperties = BaseLegacyReservationRecommendationPropertiesImpl{}

type BaseLegacyReservationRecommendationPropertiesImpl struct {
	CostWithNoReservedInstances    *float64       `json:"costWithNoReservedInstances,omitempty"`
	FirstUsageDate                 *string        `json:"firstUsageDate,omitempty"`
	InstanceFlexibilityGroup       *string        `json:"instanceFlexibilityGroup,omitempty"`
	InstanceFlexibilityRatio       *float64       `json:"instanceFlexibilityRatio,omitempty"`
	LastUsageDate                  *string        `json:"lastUsageDate,omitempty"`
	LookBackPeriod                 *string        `json:"lookBackPeriod,omitempty"`
	MeterId                        *string        `json:"meterId,omitempty"`
	NetSavings                     *float64       `json:"netSavings,omitempty"`
	NormalizedSize                 *string        `json:"normalizedSize,omitempty"`
	RecommendedQuantity            *float64       `json:"recommendedQuantity,omitempty"`
	RecommendedQuantityNormalized  *float64       `json:"recommendedQuantityNormalized,omitempty"`
	ResourceType                   *string        `json:"resourceType,omitempty"`
	Scope                          string         `json:"scope"`
	SkuProperties                  *[]SkuProperty `json:"skuProperties,omitempty"`
	Term                           *string        `json:"term,omitempty"`
	TotalCostWithReservedInstances *float64       `json:"totalCostWithReservedInstances,omitempty"`
	TotalHours                     *int64         `json:"totalHours,omitempty"`
}

func (s BaseLegacyReservationRecommendationPropertiesImpl) LegacyReservationRecommendationProperties() BaseLegacyReservationRecommendationPropertiesImpl {
	return s
}

var _ LegacyReservationRecommendationProperties = RawLegacyReservationRecommendationPropertiesImpl{}

// RawLegacyReservationRecommendationPropertiesImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawLegacyReservationRecommendationPropertiesImpl struct {
	legacyReservationRecommendationProperties BaseLegacyReservationRecommendationPropertiesImpl
	Type                                      string
	Values                                    map[string]interface{}
}

func (s RawLegacyReservationRecommendationPropertiesImpl) LegacyReservationRecommendationProperties() BaseLegacyReservationRecommendationPropertiesImpl {
	return s.legacyReservationRecommendationProperties
}

func UnmarshalLegacyReservationRecommendationPropertiesImplementation(input []byte) (LegacyReservationRecommendationProperties, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling LegacyReservationRecommendationProperties into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["scope"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Shared") {
		var out LegacySharedScopeReservationRecommendationProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LegacySharedScopeReservationRecommendationProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Single") {
		var out LegacySingleScopeReservationRecommendationProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into LegacySingleScopeReservationRecommendationProperties: %+v", err)
		}
		return out, nil
	}

	var parent BaseLegacyReservationRecommendationPropertiesImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseLegacyReservationRecommendationPropertiesImpl: %+v", err)
	}

	return RawLegacyReservationRecommendationPropertiesImpl{
		legacyReservationRecommendationProperties: parent,
		Type:   value,
		Values: temp,
	}, nil

}
