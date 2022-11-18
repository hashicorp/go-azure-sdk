package usagedetails

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ModernUsageDetailProperties struct {
	AdditionalInfo               *string           `json:"additionalInfo,omitempty"`
	BenefitId                    *string           `json:"benefitId,omitempty"`
	BenefitName                  *string           `json:"benefitName,omitempty"`
	BillingAccountId             *string           `json:"billingAccountId,omitempty"`
	BillingAccountName           *string           `json:"billingAccountName,omitempty"`
	BillingCurrencyCode          *string           `json:"billingCurrencyCode,omitempty"`
	BillingPeriodEndDate         *string           `json:"billingPeriodEndDate,omitempty"`
	BillingPeriodStartDate       *string           `json:"billingPeriodStartDate,omitempty"`
	BillingProfileId             *string           `json:"billingProfileId,omitempty"`
	BillingProfileName           *string           `json:"billingProfileName,omitempty"`
	ChargeType                   *string           `json:"chargeType,omitempty"`
	ConsumedService              *string           `json:"consumedService,omitempty"`
	CostAllocationRuleName       *string           `json:"costAllocationRuleName,omitempty"`
	CostCenter                   *string           `json:"costCenter,omitempty"`
	CostInBillingCurrency        *float64          `json:"costInBillingCurrency,omitempty"`
	CostInPricingCurrency        *float64          `json:"costInPricingCurrency,omitempty"`
	CostInUSD                    *float64          `json:"costInUSD,omitempty"`
	CustomerName                 *string           `json:"customerName,omitempty"`
	CustomerTenantId             *string           `json:"customerTenantId,omitempty"`
	Date                         *string           `json:"date,omitempty"`
	EffectivePrice               *float64          `json:"effectivePrice,omitempty"`
	ExchangeRate                 *string           `json:"exchangeRate,omitempty"`
	ExchangeRateDate             *string           `json:"exchangeRateDate,omitempty"`
	ExchangeRatePricingToBilling *float64          `json:"exchangeRatePricingToBilling,omitempty"`
	Frequency                    *string           `json:"frequency,omitempty"`
	InstanceName                 *string           `json:"instanceName,omitempty"`
	InvoiceId                    *string           `json:"invoiceId,omitempty"`
	InvoiceSectionId             *string           `json:"invoiceSectionId,omitempty"`
	InvoiceSectionName           *string           `json:"invoiceSectionName,omitempty"`
	IsAzureCreditEligible        *bool             `json:"isAzureCreditEligible,omitempty"`
	MarketPrice                  *float64          `json:"marketPrice,omitempty"`
	MeterCategory                *string           `json:"meterCategory,omitempty"`
	MeterId                      *string           `json:"meterId,omitempty"`
	MeterName                    *string           `json:"meterName,omitempty"`
	MeterRegion                  *string           `json:"meterRegion,omitempty"`
	MeterSubCategory             *string           `json:"meterSubCategory,omitempty"`
	PartnerEarnedCreditApplied   *string           `json:"partnerEarnedCreditApplied,omitempty"`
	PartnerEarnedCreditRate      *float64          `json:"partnerEarnedCreditRate,omitempty"`
	PartnerName                  *string           `json:"partnerName,omitempty"`
	PartnerTenantId              *string           `json:"partnerTenantId,omitempty"`
	PayGPrice                    *float64          `json:"payGPrice,omitempty"`
	PaygCostInBillingCurrency    *float64          `json:"paygCostInBillingCurrency,omitempty"`
	PaygCostInUSD                *float64          `json:"paygCostInUSD,omitempty"`
	PreviousInvoiceId            *string           `json:"previousInvoiceId,omitempty"`
	PricingCurrencyCode          *string           `json:"pricingCurrencyCode,omitempty"`
	PricingModel                 *PricingModelType `json:"pricingModel,omitempty"`
	Product                      *string           `json:"product,omitempty"`
	ProductIdentifier            *string           `json:"productIdentifier,omitempty"`
	ProductOrderId               *string           `json:"productOrderId,omitempty"`
	ProductOrderName             *string           `json:"productOrderName,omitempty"`
	Provider                     *string           `json:"provider,omitempty"`
	PublisherId                  *string           `json:"publisherId,omitempty"`
	PublisherName                *string           `json:"publisherName,omitempty"`
	PublisherType                *string           `json:"publisherType,omitempty"`
	Quantity                     *float64          `json:"quantity,omitempty"`
	ResellerMpnId                *string           `json:"resellerMpnId,omitempty"`
	ResellerName                 *string           `json:"resellerName,omitempty"`
	ReservationId                *string           `json:"reservationId,omitempty"`
	ReservationName              *string           `json:"reservationName,omitempty"`
	ResourceGroup                *string           `json:"resourceGroup,omitempty"`
	ResourceLocation             *string           `json:"resourceLocation,omitempty"`
	ResourceLocationNormalized   *string           `json:"resourceLocationNormalized,omitempty"`
	ServiceFamily                *string           `json:"serviceFamily,omitempty"`
	ServiceInfo1                 *string           `json:"serviceInfo1,omitempty"`
	ServiceInfo2                 *string           `json:"serviceInfo2,omitempty"`
	ServicePeriodEndDate         *string           `json:"servicePeriodEndDate,omitempty"`
	ServicePeriodStartDate       *string           `json:"servicePeriodStartDate,omitempty"`
	SubscriptionGuid             *string           `json:"subscriptionGuid,omitempty"`
	SubscriptionName             *string           `json:"subscriptionName,omitempty"`
	Term                         *string           `json:"term,omitempty"`
	UnitOfMeasure                *string           `json:"unitOfMeasure,omitempty"`
	UnitPrice                    *float64          `json:"unitPrice,omitempty"`
}

func (o *ModernUsageDetailProperties) GetBillingPeriodEndDateAsTime() (*time.Time, error) {
	if o.BillingPeriodEndDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.BillingPeriodEndDate, "2006-01-02T15:04:05Z07:00")
}

func (o *ModernUsageDetailProperties) SetBillingPeriodEndDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.BillingPeriodEndDate = &formatted
}

func (o *ModernUsageDetailProperties) GetBillingPeriodStartDateAsTime() (*time.Time, error) {
	if o.BillingPeriodStartDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.BillingPeriodStartDate, "2006-01-02T15:04:05Z07:00")
}

func (o *ModernUsageDetailProperties) SetBillingPeriodStartDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.BillingPeriodStartDate = &formatted
}

func (o *ModernUsageDetailProperties) GetDateAsTime() (*time.Time, error) {
	if o.Date == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.Date, "2006-01-02T15:04:05Z07:00")
}

func (o *ModernUsageDetailProperties) SetDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Date = &formatted
}

func (o *ModernUsageDetailProperties) GetExchangeRateDateAsTime() (*time.Time, error) {
	if o.ExchangeRateDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ExchangeRateDate, "2006-01-02T15:04:05Z07:00")
}

func (o *ModernUsageDetailProperties) SetExchangeRateDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ExchangeRateDate = &formatted
}

func (o *ModernUsageDetailProperties) GetServicePeriodEndDateAsTime() (*time.Time, error) {
	if o.ServicePeriodEndDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ServicePeriodEndDate, "2006-01-02T15:04:05Z07:00")
}

func (o *ModernUsageDetailProperties) SetServicePeriodEndDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ServicePeriodEndDate = &formatted
}

func (o *ModernUsageDetailProperties) GetServicePeriodStartDateAsTime() (*time.Time, error) {
	if o.ServicePeriodStartDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ServicePeriodStartDate, "2006-01-02T15:04:05Z07:00")
}

func (o *ModernUsageDetailProperties) SetServicePeriodStartDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.ServicePeriodStartDate = &formatted
}
