package transactions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TransactionProperties struct {
	AzureCreditApplied        *Amount              `json:"azureCreditApplied,omitempty"`
	BillingCurrency           *string              `json:"billingCurrency,omitempty"`
	BillingProfileDisplayName *string              `json:"billingProfileDisplayName,omitempty"`
	BillingProfileId          *string              `json:"billingProfileId,omitempty"`
	CustomerDisplayName       *string              `json:"customerDisplayName,omitempty"`
	CustomerId                *string              `json:"customerId,omitempty"`
	Date                      *string              `json:"date,omitempty"`
	Discount                  *float64             `json:"discount,omitempty"`
	EffectivePrice            *Amount              `json:"effectivePrice,omitempty"`
	ExchangeRate              *float64             `json:"exchangeRate,omitempty"`
	Invoice                   *string              `json:"invoice,omitempty"`
	InvoiceSectionDisplayName *string              `json:"invoiceSectionDisplayName,omitempty"`
	InvoiceSectionId          *string              `json:"invoiceSectionId,omitempty"`
	Kind                      *TransactionTypeKind `json:"kind,omitempty"`
	MarketPrice               *Amount              `json:"marketPrice,omitempty"`
	OrderId                   *string              `json:"orderId,omitempty"`
	OrderName                 *string              `json:"orderName,omitempty"`
	PricingCurrency           *string              `json:"pricingCurrency,omitempty"`
	ProductDescription        *string              `json:"productDescription,omitempty"`
	ProductFamily             *string              `json:"productFamily,omitempty"`
	ProductType               *string              `json:"productType,omitempty"`
	ProductTypeId             *string              `json:"productTypeId,omitempty"`
	Quantity                  *int64               `json:"quantity,omitempty"`
	ServicePeriodEndDate      *string              `json:"servicePeriodEndDate,omitempty"`
	ServicePeriodStartDate    *string              `json:"servicePeriodStartDate,omitempty"`
	SubTotal                  *Amount              `json:"subTotal,omitempty"`
	SubscriptionId            *string              `json:"subscriptionId,omitempty"`
	SubscriptionName          *string              `json:"subscriptionName,omitempty"`
	Tax                       *Amount              `json:"tax,omitempty"`
	TransactionAmount         *Amount              `json:"transactionAmount,omitempty"`
	TransactionType           *ReservationType     `json:"transactionType,omitempty"`
	UnitOfMeasure             *string              `json:"unitOfMeasure,omitempty"`
	UnitType                  *string              `json:"unitType,omitempty"`
	Units                     *float64             `json:"units,omitempty"`
}
