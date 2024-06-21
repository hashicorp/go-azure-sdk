package lots

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LotProperties struct {
	BillingCurrency                 *string                 `json:"billingCurrency,omitempty"`
	ClosedBalance                   *Amount                 `json:"closedBalance,omitempty"`
	ClosedBalanceInBillingCurrency  *AmountWithExchangeRate `json:"closedBalanceInBillingCurrency,omitempty"`
	CreditCurrency                  *string                 `json:"creditCurrency,omitempty"`
	ETag                            *string                 `json:"eTag,omitempty"`
	ExpirationDate                  *string                 `json:"expirationDate,omitempty"`
	IsEstimatedBalance              *bool                   `json:"isEstimatedBalance,omitempty"`
	OrganizationType                *OrganizationType       `json:"OrganizationType,omitempty"`
	OriginalAmount                  *Amount                 `json:"originalAmount,omitempty"`
	OriginalAmountInBillingCurrency *AmountWithExchangeRate `json:"originalAmountInBillingCurrency,omitempty"`
	PoNumber                        *string                 `json:"poNumber,omitempty"`
	PurchasedDate                   *string                 `json:"purchasedDate,omitempty"`
	Reseller                        *Reseller               `json:"reseller,omitempty"`
	Source                          *LotSource              `json:"source,omitempty"`
	StartDate                       *string                 `json:"startDate,omitempty"`
	Status                          *Status                 `json:"status,omitempty"`
	UsedAmount                      *Amount                 `json:"usedAmount,omitempty"`
}
