package credits

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CreditSummaryProperties struct {
	BalanceSummary           *CreditBalanceSummary `json:"balanceSummary"`
	BillingCurrency          *string               `json:"billingCurrency,omitempty"`
	CreditCurrency           *string               `json:"creditCurrency,omitempty"`
	ExpiredCredit            *Amount               `json:"expiredCredit"`
	PendingCreditAdjustments *Amount               `json:"pendingCreditAdjustments"`
	PendingEligibleCharges   *Amount               `json:"pendingEligibleCharges"`
	Reseller                 *Reseller             `json:"reseller"`
}
