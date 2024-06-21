package events

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventProperties struct {
	Adjustments                    *Amount                 `json:"adjustments,omitempty"`
	AdjustmentsInBillingCurrency   *AmountWithExchangeRate `json:"adjustmentsInBillingCurrency,omitempty"`
	BillingAccountDisplayName      *string                 `json:"billingAccountDisplayName,omitempty"`
	BillingAccountId               *string                 `json:"billingAccountId,omitempty"`
	BillingCurrency                *string                 `json:"billingCurrency,omitempty"`
	BillingProfileDisplayName      *string                 `json:"billingProfileDisplayName,omitempty"`
	BillingProfileId               *string                 `json:"billingProfileId,omitempty"`
	CanceledCredit                 *Amount                 `json:"canceledCredit,omitempty"`
	Charges                        *Amount                 `json:"charges,omitempty"`
	ChargesInBillingCurrency       *AmountWithExchangeRate `json:"chargesInBillingCurrency,omitempty"`
	ClosedBalance                  *Amount                 `json:"closedBalance,omitempty"`
	ClosedBalanceInBillingCurrency *AmountWithExchangeRate `json:"closedBalanceInBillingCurrency,omitempty"`
	CreditCurrency                 *string                 `json:"creditCurrency,omitempty"`
	CreditExpired                  *Amount                 `json:"creditExpired,omitempty"`
	CreditExpiredInBillingCurrency *AmountWithExchangeRate `json:"creditExpiredInBillingCurrency,omitempty"`
	Description                    *string                 `json:"description,omitempty"`
	ETag                           *string                 `json:"eTag,omitempty"`
	EventType                      *EventType              `json:"eventType,omitempty"`
	InvoiceNumber                  *string                 `json:"invoiceNumber,omitempty"`
	IsEstimatedBalance             *bool                   `json:"isEstimatedBalance,omitempty"`
	LotId                          *string                 `json:"lotId,omitempty"`
	LotSource                      *string                 `json:"lotSource,omitempty"`
	NewCredit                      *Amount                 `json:"newCredit,omitempty"`
	NewCreditInBillingCurrency     *AmountWithExchangeRate `json:"newCreditInBillingCurrency,omitempty"`
	Reseller                       *Reseller               `json:"reseller,omitempty"`
	TransactionDate                *string                 `json:"transactionDate,omitempty"`
}
