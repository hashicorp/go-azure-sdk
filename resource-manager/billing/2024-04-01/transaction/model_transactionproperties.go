package transaction

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type TransactionProperties struct {
	AzureCreditApplied               *Amount                   `json:"azureCreditApplied,omitempty"`
	AzurePlan                        *string                   `json:"azurePlan,omitempty"`
	BillingCurrency                  *string                   `json:"billingCurrency,omitempty"`
	BillingProfileDisplayName        *interface{}              `json:"billingProfileDisplayName,omitempty"`
	BillingProfileId                 *string                   `json:"billingProfileId,omitempty"`
	ConsumptionCommitmentDecremented *Amount                   `json:"consumptionCommitmentDecremented,omitempty"`
	CreditType                       *CreditType               `json:"creditType,omitempty"`
	CustomerDisplayName              *string                   `json:"customerDisplayName,omitempty"`
	CustomerId                       *string                   `json:"customerId,omitempty"`
	Date                             *string                   `json:"date,omitempty"`
	Discount                         *float64                  `json:"discount,omitempty"`
	EffectivePrice                   *Amount                   `json:"effectivePrice,omitempty"`
	ExchangeRate                     *float64                  `json:"exchangeRate,omitempty"`
	Invoice                          *string                   `json:"invoice,omitempty"`
	InvoiceId                        *string                   `json:"invoiceId,omitempty"`
	InvoiceSectionDisplayName        *string                   `json:"invoiceSectionDisplayName,omitempty"`
	InvoiceSectionId                 *string                   `json:"invoiceSectionId,omitempty"`
	IsThirdParty                     *bool                     `json:"isThirdParty,omitempty"`
	Kind                             *TransactionKind          `json:"kind,omitempty"`
	MarketPrice                      *Amount                   `json:"marketPrice,omitempty"`
	PartNumber                       *string                   `json:"partNumber,omitempty"`
	PricingCurrency                  *string                   `json:"pricingCurrency,omitempty"`
	ProductDescription               *string                   `json:"productDescription,omitempty"`
	ProductFamily                    *string                   `json:"productFamily,omitempty"`
	ProductType                      *string                   `json:"productType,omitempty"`
	ProductTypeId                    *string                   `json:"productTypeId,omitempty"`
	Quantity                         *int64                    `json:"quantity,omitempty"`
	ReasonCode                       *string                   `json:"reasonCode,omitempty"`
	RefundTransactionDetails         *RefundTransactionDetails `json:"refundTransactionDetails,omitempty"`
	ServicePeriodEndDate             *string                   `json:"servicePeriodEndDate,omitempty"`
	ServicePeriodStartDate           *string                   `json:"servicePeriodStartDate,omitempty"`
	SpecialTaxationType              *SpecialTaxationType      `json:"specialTaxationType,omitempty"`
	SubTotal                         *Amount                   `json:"subTotal,omitempty"`
	Tax                              *Amount                   `json:"tax,omitempty"`
	TransactionAmount                *Amount                   `json:"transactionAmount,omitempty"`
	TransactionType                  *string                   `json:"transactionType,omitempty"`
	UnitOfMeasure                    *string                   `json:"unitOfMeasure,omitempty"`
	UnitType                         *string                   `json:"unitType,omitempty"`
	Units                            *float64                  `json:"units,omitempty"`
}

func (o *TransactionProperties) GetDateAsTime() (*time.Time, error) {
	if o.Date == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.Date, "2006-01-02T15:04:05Z07:00")
}

func (o *TransactionProperties) SetDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Date = &formatted
}

func (o *TransactionProperties) GetServicePeriodEndDateAsTime() (*time.Time, error) {
	if o.ServicePeriodEndDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ServicePeriodEndDate, "2006-01-02T15:04:05Z07:00")
}

func (o *TransactionProperties) SetServicePeriodEndDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ServicePeriodEndDate = &formatted
}

func (o *TransactionProperties) GetServicePeriodStartDateAsTime() (*time.Time, error) {
	if o.ServicePeriodStartDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ServicePeriodStartDate, "2006-01-02T15:04:05Z07:00")
}

func (o *TransactionProperties) SetServicePeriodStartDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ServicePeriodStartDate = &formatted
}
