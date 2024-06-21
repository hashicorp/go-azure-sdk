package benefitrecommendations

import (
	"encoding/json"
	"fmt"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BenefitRecommendationProperties = SharedScopeBenefitRecommendationProperties{}

type SharedScopeBenefitRecommendationProperties struct {

	// Fields inherited from BenefitRecommendationProperties
	AllRecommendationDetails *AllSavingsList             `json:"allRecommendationDetails,omitempty"`
	ArmSkuName               *string                     `json:"armSkuName,omitempty"`
	CommitmentGranularity    *Grain                      `json:"commitmentGranularity,omitempty"`
	CostWithoutBenefit       *float64                    `json:"costWithoutBenefit,omitempty"`
	CurrencyCode             *string                     `json:"currencyCode,omitempty"`
	FirstConsumptionDate     *string                     `json:"firstConsumptionDate,omitempty"`
	LastConsumptionDate      *string                     `json:"lastConsumptionDate,omitempty"`
	LookBackPeriod           *LookBackPeriod             `json:"lookBackPeriod,omitempty"`
	RecommendationDetails    *AllSavingsBenefitDetails   `json:"recommendationDetails,omitempty"`
	Term                     *Term                       `json:"term,omitempty"`
	TotalHours               *int64                      `json:"totalHours,omitempty"`
	Usage                    *RecommendationUsageDetails `json:"usage,omitempty"`
}

var _ json.Marshaler = SharedScopeBenefitRecommendationProperties{}

func (s SharedScopeBenefitRecommendationProperties) MarshalJSON() ([]byte, error) {
	type wrapper SharedScopeBenefitRecommendationProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SharedScopeBenefitRecommendationProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SharedScopeBenefitRecommendationProperties: %+v", err)
	}
	decoded["scope"] = "Shared"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SharedScopeBenefitRecommendationProperties: %+v", err)
	}

	return encoded, nil
}
