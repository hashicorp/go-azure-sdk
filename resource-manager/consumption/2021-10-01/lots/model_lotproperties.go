package lots

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LotProperties struct {
	BillingCurrency                 *string                 `json:"billingCurrency,omitempty"`
	ClosedBalance                   *Amount                 `json:"closedBalance"`
	ClosedBalanceInBillingCurrency  *AmountWithExchangeRate `json:"closedBalanceInBillingCurrency"`
	CreditCurrency                  *string                 `json:"creditCurrency,omitempty"`
	ETag                            *string                 `json:"eTag,omitempty"`
	ExpirationDate                  *string                 `json:"expirationDate,omitempty"`
	OriginalAmount                  *Amount                 `json:"originalAmount"`
	OriginalAmountInBillingCurrency *AmountWithExchangeRate `json:"originalAmountInBillingCurrency"`
	PoNumber                        *string                 `json:"poNumber,omitempty"`
	PurchasedDate                   *string                 `json:"purchasedDate,omitempty"`
	Reseller                        *Reseller               `json:"reseller"`
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

func (o *LotProperties) SetExpirationDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ExpirationDate = &formatted
}

func (o *LotProperties) GetPurchasedDateAsTime() (*time.Time, error) {
	if o.PurchasedDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.PurchasedDate, "2006-01-02T15:04:05Z07:00")
}

func (o *LotProperties) SetPurchasedDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.PurchasedDate = &formatted
}

func (o *LotProperties) GetStartDateAsTime() (*time.Time, error) {
	if o.StartDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.StartDate, "2006-01-02T15:04:05Z07:00")
}

func (o *LotProperties) SetStartDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.StartDate = &formatted
}
