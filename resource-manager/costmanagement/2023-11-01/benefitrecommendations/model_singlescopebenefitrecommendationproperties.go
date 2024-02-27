package benefitrecommendations

import (
	"encoding/json"
	"fmt"
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

var _ BenefitRecommendationProperties = SingleScopeBenefitRecommendationProperties{}

type SingleScopeBenefitRecommendationProperties struct {
	ResourceGroup  *string `json:"resourceGroup,omitempty"`
	SubscriptionId *string `json:"subscriptionId,omitempty"`

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

func (o *SingleScopeBenefitRecommendationProperties) GetFirstConsumptionDateAsTime() (*time.Time, error) {
	if o.FirstConsumptionDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.FirstConsumptionDate, "2006-01-02T15:04:05Z07:00")
}

func (o *SingleScopeBenefitRecommendationProperties) GetLastConsumptionDateAsTime() (*time.Time, error) {
	if o.LastConsumptionDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastConsumptionDate, "2006-01-02T15:04:05Z07:00")
}

var _ json.Marshaler = SingleScopeBenefitRecommendationProperties{}

func (s SingleScopeBenefitRecommendationProperties) MarshalJSON() ([]byte, error) {
	type wrapper SingleScopeBenefitRecommendationProperties
	wrapped := wrapper(s)
	encoded, err := json.Marshal(wrapped)
	if err != nil {
		return nil, fmt.Errorf("marshaling SingleScopeBenefitRecommendationProperties: %+v", err)
	}

	var decoded map[string]interface{}
	if err := json.Unmarshal(encoded, &decoded); err != nil {
		return nil, fmt.Errorf("unmarshaling SingleScopeBenefitRecommendationProperties: %+v", err)
	}
	decoded["scope"] = "Single"

	encoded, err = json.Marshal(decoded)
	if err != nil {
		return nil, fmt.Errorf("re-marshaling SingleScopeBenefitRecommendationProperties: %+v", err)
	}

	return encoded, nil
}
