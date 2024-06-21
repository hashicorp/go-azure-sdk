package firewalls

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PlanData struct {
	BillingCycle  BillingCycle `json:"billingCycle"`
	EffectiveDate *string      `json:"effectiveDate,omitempty"`
	PlanId        string       `json:"planId"`
	UsageType     *UsageType   `json:"usageType,omitempty"`
}
