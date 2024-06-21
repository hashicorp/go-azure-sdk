package benefitutilizationsummaries

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SavingsPlanUtilizationSummaryProperties struct {
	ArmSkuName               *string      `json:"armSkuName,omitempty"`
	AvgUtilizationPercentage *float64     `json:"avgUtilizationPercentage,omitempty"`
	BenefitId                *string      `json:"benefitId,omitempty"`
	BenefitOrderId           *string      `json:"benefitOrderId,omitempty"`
	BenefitType              *BenefitKind `json:"benefitType,omitempty"`
	MaxUtilizationPercentage *float64     `json:"maxUtilizationPercentage,omitempty"`
	MinUtilizationPercentage *float64     `json:"minUtilizationPercentage,omitempty"`
	UsageDate                *string      `json:"usageDate,omitempty"`
}
