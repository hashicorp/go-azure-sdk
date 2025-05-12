package benefitrecommendations

import (
	"encoding/json"
	"fmt"
	"strings"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BenefitRecommendationProperties interface {
	BenefitRecommendationProperties() BaseBenefitRecommendationPropertiesImpl
}

var _ BenefitRecommendationProperties = BaseBenefitRecommendationPropertiesImpl{}

type BaseBenefitRecommendationPropertiesImpl struct {
	AllRecommendationDetails *AllSavingsList             `json:"allRecommendationDetails,omitempty"`
	ArmSkuName               *string                     `json:"armSkuName,omitempty"`
	CommitmentGranularity    *Grain                      `json:"commitmentGranularity,omitempty"`
	CostWithoutBenefit       *float64                    `json:"costWithoutBenefit,omitempty"`
	CurrencyCode             *string                     `json:"currencyCode,omitempty"`
	FirstConsumptionDate     *string                     `json:"firstConsumptionDate,omitempty"`
	LastConsumptionDate      *string                     `json:"lastConsumptionDate,omitempty"`
	LookBackPeriod           *LookBackPeriod             `json:"lookBackPeriod,omitempty"`
	RecommendationDetails    *AllSavingsBenefitDetails   `json:"recommendationDetails,omitempty"`
	Scope                    Scope                       `json:"scope"`
	Term                     *Term                       `json:"term,omitempty"`
	TotalHours               *int64                      `json:"totalHours,omitempty"`
	Usage                    *RecommendationUsageDetails `json:"usage,omitempty"`
}

func (s BaseBenefitRecommendationPropertiesImpl) BenefitRecommendationProperties() BaseBenefitRecommendationPropertiesImpl {
	return s
}

var _ BenefitRecommendationProperties = RawBenefitRecommendationPropertiesImpl{}

// RawBenefitRecommendationPropertiesImpl is returned when the Discriminated Value doesn't match any of the defined types
// NOTE: this should only be used when a type isn't defined for this type of Object (as a workaround)
// and is used only for Deserialization (e.g. this cannot be used as a Request Payload).
type RawBenefitRecommendationPropertiesImpl struct {
	benefitRecommendationProperties BaseBenefitRecommendationPropertiesImpl
	Type                            string
	Values                          map[string]interface{}
}

func (s RawBenefitRecommendationPropertiesImpl) BenefitRecommendationProperties() BaseBenefitRecommendationPropertiesImpl {
	return s.benefitRecommendationProperties
}

func UnmarshalBenefitRecommendationPropertiesImplementation(input []byte) (BenefitRecommendationProperties, error) {
	if input == nil {
		return nil, nil
	}

	var temp map[string]interface{}
	if err := json.Unmarshal(input, &temp); err != nil {
		return nil, fmt.Errorf("unmarshaling BenefitRecommendationProperties into map[string]interface: %+v", err)
	}

	var value string
	if v, ok := temp["scope"]; ok {
		value = fmt.Sprintf("%v", v)
	}

	if strings.EqualFold(value, "Shared") {
		var out SharedScopeBenefitRecommendationProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SharedScopeBenefitRecommendationProperties: %+v", err)
		}
		return out, nil
	}

	if strings.EqualFold(value, "Single") {
		var out SingleScopeBenefitRecommendationProperties
		if err := json.Unmarshal(input, &out); err != nil {
			return nil, fmt.Errorf("unmarshaling into SingleScopeBenefitRecommendationProperties: %+v", err)
		}
		return out, nil
	}

	var parent BaseBenefitRecommendationPropertiesImpl
	if err := json.Unmarshal(input, &parent); err != nil {
		return nil, fmt.Errorf("unmarshaling into BaseBenefitRecommendationPropertiesImpl: %+v", err)
	}

	return RawBenefitRecommendationPropertiesImpl{
		benefitRecommendationProperties: parent,
		Type:                            value,
		Values:                          temp,
	}, nil

}
