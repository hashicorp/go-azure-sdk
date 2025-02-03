package pricesheets

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MCAPriceSheetProperties struct {
	BasePrice          *string  `json:"basePrice,omitempty"`
	BillingAccountID   *string  `json:"billingAccountID,omitempty"`
	BillingAccountName *string  `json:"billingAccountName,omitempty"`
	BillingCurrency    *string  `json:"billingCurrency,omitempty"`
	BillingProfileId   *string  `json:"billingProfileId,omitempty"`
	BillingProfileName *string  `json:"billingProfileName,omitempty"`
	Currency           *string  `json:"currency,omitempty"`
	EffectiveEndDate   *string  `json:"effectiveEndDate,omitempty"`
	EffectiveStartDate *string  `json:"effectiveStartDate,omitempty"`
	MarketPrice        *string  `json:"marketPrice,omitempty"`
	MeterCategory      *string  `json:"meterCategory,omitempty"`
	MeterId            *string  `json:"meterId,omitempty"`
	MeterName          *string  `json:"meterName,omitempty"`
	MeterRegion        *string  `json:"meterRegion,omitempty"`
	MeterSubCategory   *string  `json:"meterSubCategory,omitempty"`
	MeterType          *string  `json:"meterType,omitempty"`
	PriceType          *string  `json:"priceType,omitempty"`
	Product            *string  `json:"product,omitempty"`
	ProductId          *string  `json:"productId,omitempty"`
	ServiceFamily      *float64 `json:"serviceFamily,omitempty"`
	SkuId              *string  `json:"skuId,omitempty"`
	Term               *string  `json:"term,omitempty"`
	TierMinimumUnits   *string  `json:"tierMinimumUnits,omitempty"`
	UnitOfMeasure      *string  `json:"unitOfMeasure,omitempty"`
	UnitPrice          *string  `json:"unitPrice,omitempty"`
}

func (o *MCAPriceSheetProperties) GetEffectiveEndDateAsTime() (*time.Time, error) {
	if o.EffectiveEndDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EffectiveEndDate, "2006-01-02T15:04:05Z07:00")
}

func (o *MCAPriceSheetProperties) SetEffectiveEndDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EffectiveEndDate = &formatted
}

func (o *MCAPriceSheetProperties) GetEffectiveStartDateAsTime() (*time.Time, error) {
	if o.EffectiveStartDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EffectiveStartDate, "2006-01-02T15:04:05Z07:00")
}

func (o *MCAPriceSheetProperties) SetEffectiveStartDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EffectiveStartDate = &formatted
}
