package benefitutilizationsummariesasync

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BenefitUtilizationSummariesRequest struct {
	BenefitId        *string      `json:"benefitId,omitempty"`
	BenefitOrderId   *string      `json:"benefitOrderId,omitempty"`
	BillingAccountId *string      `json:"billingAccountId,omitempty"`
	BillingProfileId *string      `json:"billingProfileId,omitempty"`
	EndDate          string       `json:"endDate"`
	Grain            Grain        `json:"grain"`
	Kind             *BenefitKind `json:"kind,omitempty"`
	StartDate        string       `json:"startDate"`
}
