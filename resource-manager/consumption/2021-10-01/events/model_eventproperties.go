package events

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EventProperties struct {
	Adjustments                    *Amount                 `json:"adjustments"`
	AdjustmentsInBillingCurrency   *AmountWithExchangeRate `json:"adjustmentsInBillingCurrency"`
	BillingCurrency                *string                 `json:"billingCurrency,omitempty"`
	BillingProfileDisplayName      *string                 `json:"billingProfileDisplayName,omitempty"`
	BillingProfileId               *string                 `json:"billingProfileId,omitempty"`
	CanceledCredit                 *Amount                 `json:"canceledCredit"`
	Charges                        *Amount                 `json:"charges"`
	ChargesInBillingCurrency       *AmountWithExchangeRate `json:"chargesInBillingCurrency"`
	ClosedBalance                  *Amount                 `json:"closedBalance"`
	ClosedBalanceInBillingCurrency *AmountWithExchangeRate `json:"closedBalanceInBillingCurrency"`
	CreditCurrency                 *string                 `json:"creditCurrency,omitempty"`
	CreditExpired                  *Amount                 `json:"creditExpired"`
	CreditExpiredInBillingCurrency *AmountWithExchangeRate `json:"creditExpiredInBillingCurrency"`
	Description                    *string                 `json:"description,omitempty"`
	ETag                           *string                 `json:"eTag,omitempty"`
	EventType                      *EventType              `json:"eventType,omitempty"`
	InvoiceNumber                  *string                 `json:"invoiceNumber,omitempty"`
	LotId                          *string                 `json:"lotId,omitempty"`
	LotSource                      *string                 `json:"lotSource,omitempty"`
	NewCredit                      *Amount                 `json:"newCredit"`
	NewCreditInBillingCurrency     *AmountWithExchangeRate `json:"newCreditInBillingCurrency"`
	Reseller                       *Reseller               `json:"reseller"`
	TransactionDate                *string                 `json:"transactionDate,omitempty"`
}

func (o *EventProperties) GetTransactionDateAsTime() (*time.Time, error) {
	if o.TransactionDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.TransactionDate, "2006-01-02T15:04:05Z07:00")
}

func (o *EventProperties) SetTransactionDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.TransactionDate = &formatted
}
