package pricesheet20220601

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PriceSheetPropertiesV2 struct {
	BillingPeriodId  *string         `json:"billingPeriodId,omitempty"`
	CurrencyCode     *string         `json:"currencyCode,omitempty"`
	IncludedQuantity *float64        `json:"includedQuantity,omitempty"`
	MeterDetails     *MeterDetailsV2 `json:"meterDetails,omitempty"`
	MeterId          *string         `json:"meterId,omitempty"`
	OfferId          *string         `json:"offerId,omitempty"`
	PartNumber       *string         `json:"partNumber,omitempty"`
	UnitOfMeasure    *string         `json:"unitOfMeasure,omitempty"`
	UnitPrice        *float64        `json:"unitPrice,omitempty"`
}
