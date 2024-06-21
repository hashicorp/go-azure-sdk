package plan

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlanData struct {
	BillingCycle  *BillingCycle `json:"billingCycle,omitempty"`
	EffectiveDate *string       `json:"effectiveDate,omitempty"`
	PlanDetails   *string       `json:"planDetails,omitempty"`
	UsageType     *UsageType    `json:"usageType,omitempty"`
}
