package costs

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LabCostProperties struct {
	CreatedDate       *string                      `json:"createdDate,omitempty"`
	CurrencyCode      *string                      `json:"currencyCode,omitempty"`
	EndDateTime       *string                      `json:"endDateTime,omitempty"`
	LabCostDetails    *[]LabCostDetailsProperties  `json:"labCostDetails,omitempty"`
	LabCostSummary    *LabCostSummaryProperties    `json:"labCostSummary,omitempty"`
	ProvisioningState *string                      `json:"provisioningState,omitempty"`
	ResourceCosts     *[]LabResourceCostProperties `json:"resourceCosts,omitempty"`
	StartDateTime     *string                      `json:"startDateTime,omitempty"`
	TargetCost        *TargetCostProperties        `json:"targetCost,omitempty"`
	UniqueIdentifier  *string                      `json:"uniqueIdentifier,omitempty"`
}
