package lots

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

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
	OriginalAmount                  *Amount                 `json:"originalAmount,omitempty"`
	OriginalAmountInBillingCurrency *AmountWithExchangeRate `json:"originalAmountInBillingCurrency,omitempty"`
	PoNumber                        *string                 `json:"poNumber,omitempty"`
	PurchasedDate                   *string                 `json:"purchasedDate,omitempty"`
	Reseller                        *Reseller               `json:"reseller,omitempty"`
	Source                          *LotSource              `json:"source,omitempty"`
	StartDate                       *string                 `json:"startDate,omitempty"`
	Status                          *Status                 `json:"status,omitempty"`
}

func (o *LotProperties) GetExpirationDateAsTime() (*time.Time, error) {
	if o.ExpirationDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ExpirationDate, "2006-01-02T15:04:05Z07:00")
}

func (o *LotProperties) GetPurchasedDateAsTime() (*time.Time, error) {
	if o.PurchasedDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.PurchasedDate, "2006-01-02T15:04:05Z07:00")
}

func (o *LotProperties) GetStartDateAsTime() (*time.Time, error) {
	if o.StartDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartDate, "2006-01-02T15:04:05Z07:00")
}
