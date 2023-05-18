package benefitrecommendations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AllSavingsBenefitDetails struct {
	AverageUtilizationPercentage *float64 `json:"averageUtilizationPercentage,omitempty"`
	BenefitCost                  *float64 `json:"benefitCost,omitempty"`
	CommitmentAmount             *float64 `json:"commitmentAmount,omitempty"`
	CoveragePercentage           *float64 `json:"coveragePercentage,omitempty"`
	OverageCost                  *float64 `json:"overageCost,omitempty"`
	SavingsAmount                *float64 `json:"savingsAmount,omitempty"`
	SavingsPercentage            *float64 `json:"savingsPercentage,omitempty"`
	TotalCost                    *float64 `json:"totalCost,omitempty"`
	WastageCost                  *float64 `json:"wastageCost,omitempty"`
}
