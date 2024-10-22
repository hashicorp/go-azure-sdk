package pricesheets

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EAPriceSheetProperties struct {
	BasePrice          *string  `json:"basePrice,omitempty"`
	CurrencyCode       *string  `json:"currencyCode,omitempty"`
	EffectiveEndDate   *string  `json:"effectiveEndDate,omitempty"`
	EffectiveStartDate *string  `json:"effectiveStartDate,omitempty"`
	EnrollmentNumber   *string  `json:"enrollmentNumber,omitempty"`
	IncludedQuantity   *string  `json:"includedQuantity,omitempty"`
	MarketPrice        *string  `json:"marketPrice,omitempty"`
	MeterCategory      *string  `json:"meterCategory,omitempty"`
	MeterId            *string  `json:"meterId,omitempty"`
	MeterName          *string  `json:"meterName,omitempty"`
	MeterRegion        *string  `json:"meterRegion,omitempty"`
	MeterSubCategory   *string  `json:"meterSubCategory,omitempty"`
	MeterType          *string  `json:"meterType,omitempty"`
	OfferId            *string  `json:"offerId,omitempty"`
	PartNumber         *string  `json:"partNumber,omitempty"`
	PriceType          *string  `json:"priceType,omitempty"`
	Product            *string  `json:"product,omitempty"`
	ProductId          *string  `json:"productId,omitempty"`
	ServiceFamily      *float64 `json:"serviceFamily,omitempty"`
	SkuId              *string  `json:"skuId,omitempty"`
	Term               *string  `json:"term,omitempty"`
	UnitOfMeasure      *string  `json:"unitOfMeasure,omitempty"`
	UnitPrice          *string  `json:"unitPrice,omitempty"`
}

func (o *EAPriceSheetProperties) GetEffectiveEndDateAsTime() (*time.Time, error) {
	if o.EffectiveEndDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EffectiveEndDate, "2006-01-02T15:04:05Z07:00")
}

func (o *EAPriceSheetProperties) SetEffectiveEndDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EffectiveEndDate = &formatted
}

func (o *EAPriceSheetProperties) GetEffectiveStartDateAsTime() (*time.Time, error) {
	if o.EffectiveStartDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EffectiveStartDate, "2006-01-02T15:04:05Z07:00")
}

func (o *EAPriceSheetProperties) SetEffectiveStartDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.EffectiveStartDate = &formatted
}
