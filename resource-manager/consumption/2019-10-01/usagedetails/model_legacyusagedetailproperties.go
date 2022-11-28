package usagedetails

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LegacyUsageDetailProperties struct {
	AccountName            *string               `json:"accountName,omitempty"`
	AccountOwnerId         *string               `json:"accountOwnerId,omitempty"`
	AdditionalInfo         *string               `json:"additionalInfo,omitempty"`
	BillingAccountId       *string               `json:"billingAccountId,omitempty"`
	BillingAccountName     *string               `json:"billingAccountName,omitempty"`
	BillingCurrency        *string               `json:"billingCurrency,omitempty"`
	BillingPeriodEndDate   *string               `json:"billingPeriodEndDate,omitempty"`
	BillingPeriodStartDate *string               `json:"billingPeriodStartDate,omitempty"`
	BillingProfileId       *string               `json:"billingProfileId,omitempty"`
	BillingProfileName     *string               `json:"billingProfileName,omitempty"`
	ChargeType             *string               `json:"chargeType,omitempty"`
	ConsumedService        *string               `json:"consumedService,omitempty"`
	Cost                   *float64              `json:"cost,omitempty"`
	CostCenter             *string               `json:"costCenter,omitempty"`
	Date                   *string               `json:"date,omitempty"`
	EffectivePrice         *float64              `json:"effectivePrice,omitempty"`
	Frequency              *string               `json:"frequency,omitempty"`
	InvoiceSection         *string               `json:"invoiceSection,omitempty"`
	IsAzureCreditEligible  *bool                 `json:"isAzureCreditEligible,omitempty"`
	MeterDetails           *MeterDetailsResponse `json:"meterDetails"`
	MeterId                *string               `json:"meterId,omitempty"`
	OfferId                *string               `json:"offerId,omitempty"`
	PartNumber             *string               `json:"partNumber,omitempty"`
	PlanName               *string               `json:"planName,omitempty"`
	Product                *string               `json:"product,omitempty"`
	ProductOrderId         *string               `json:"productOrderId,omitempty"`
	ProductOrderName       *string               `json:"productOrderName,omitempty"`
	PublisherName          *string               `json:"publisherName,omitempty"`
	PublisherType          *string               `json:"publisherType,omitempty"`
	Quantity               *float64              `json:"quantity,omitempty"`
	ReservationId          *string               `json:"reservationId,omitempty"`
	ReservationName        *string               `json:"reservationName,omitempty"`
	ResourceGroup          *string               `json:"resourceGroup,omitempty"`
	ResourceId             *string               `json:"resourceId,omitempty"`
	ResourceLocation       *string               `json:"resourceLocation,omitempty"`
	ResourceName           *string               `json:"resourceName,omitempty"`
	ServiceInfo1           *string               `json:"serviceInfo1,omitempty"`
	ServiceInfo2           *string               `json:"serviceInfo2,omitempty"`
	SubscriptionId         *string               `json:"subscriptionId,omitempty"`
	SubscriptionName       *string               `json:"subscriptionName,omitempty"`
	Term                   *string               `json:"term,omitempty"`
	UnitPrice              *float64              `json:"unitPrice,omitempty"`
}

func (o *LegacyUsageDetailProperties) GetBillingPeriodEndDateAsTime() (*time.Time, error) {
	if o.BillingPeriodEndDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.BillingPeriodEndDate, "2006-01-02T15:04:05Z07:00")
}

func (o *LegacyUsageDetailProperties) SetBillingPeriodEndDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.BillingPeriodEndDate = &formatted
}

func (o *LegacyUsageDetailProperties) GetBillingPeriodStartDateAsTime() (*time.Time, error) {
	if o.BillingPeriodStartDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.BillingPeriodStartDate, "2006-01-02T15:04:05Z07:00")
}

func (o *LegacyUsageDetailProperties) SetBillingPeriodStartDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.BillingPeriodStartDate = &formatted
}

func (o *LegacyUsageDetailProperties) GetDateAsTime() (*time.Time, error) {
	if o.Date == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.Date, "2006-01-02T15:04:05Z07:00")
}

func (o *LegacyUsageDetailProperties) SetDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.Date = &formatted
}
