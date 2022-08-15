package credits

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreditSummaryProperties struct {
	BalanceSummary           *CreditBalanceSummary `json:"balanceSummary,omitempty"`
	BillingCurrency          *string               `json:"billingCurrency,omitempty"`
	CreditCurrency           *string               `json:"creditCurrency,omitempty"`
	ExpiredCredit            *Amount               `json:"expiredCredit,omitempty"`
	PendingCreditAdjustments *Amount               `json:"pendingCreditAdjustments,omitempty"`
	PendingEligibleCharges   *Amount               `json:"pendingEligibleCharges,omitempty"`
	Reseller                 *Reseller             `json:"reseller,omitempty"`
}
