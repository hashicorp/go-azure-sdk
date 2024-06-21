package benefitutilizationsummaries

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IncludedQuantityUtilizationSummaryProperties struct {
	ArmSkuName            *string      `json:"armSkuName,omitempty"`
	BenefitId             *string      `json:"benefitId,omitempty"`
	BenefitOrderId        *string      `json:"benefitOrderId,omitempty"`
	BenefitType           *BenefitKind `json:"benefitType,omitempty"`
	UsageDate             *string      `json:"usageDate,omitempty"`
	UtilizationPercentage *float64     `json:"utilizationPercentage,omitempty"`
}
