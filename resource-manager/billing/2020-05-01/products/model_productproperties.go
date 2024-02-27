package products

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ProductProperties struct {
	AutoRenew                 *AutoRenew         `json:"autoRenew,omitempty"`
	AvailabilityId            *string            `json:"availabilityId,omitempty"`
	BillingFrequency          *BillingFrequency  `json:"billingFrequency,omitempty"`
	BillingProfileDisplayName *string            `json:"billingProfileDisplayName,omitempty"`
	BillingProfileId          *string            `json:"billingProfileId,omitempty"`
	CustomerDisplayName       *string            `json:"customerDisplayName,omitempty"`
	CustomerId                *string            `json:"customerId,omitempty"`
	DisplayName               *string            `json:"displayName,omitempty"`
	EndDate                   *string            `json:"endDate,omitempty"`
	InvoiceSectionDisplayName *string            `json:"invoiceSectionDisplayName,omitempty"`
	InvoiceSectionId          *string            `json:"invoiceSectionId,omitempty"`
	LastCharge                *Amount            `json:"lastCharge,omitempty"`
	LastChargeDate            *string            `json:"lastChargeDate,omitempty"`
	ProductType               *string            `json:"productType,omitempty"`
	ProductTypeId             *string            `json:"productTypeId,omitempty"`
	PurchaseDate              *string            `json:"purchaseDate,omitempty"`
	Quantity                  *float64           `json:"quantity,omitempty"`
	Reseller                  *Reseller          `json:"reseller,omitempty"`
	SkuDescription            *string            `json:"skuDescription,omitempty"`
	SkuId                     *string            `json:"skuId,omitempty"`
	Status                    *ProductStatusType `json:"status,omitempty"`
	TenantId                  *string            `json:"tenantId,omitempty"`
}

func (o *ProductProperties) GetEndDateAsTime() (*time.Time, error) {
	if o.EndDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.EndDate, "2006-01-02T15:04:05Z07:00")
}

func (o *ProductProperties) GetLastChargeDateAsTime() (*time.Time, error) {
	if o.LastChargeDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastChargeDate, "2006-01-02T15:04:05Z07:00")
}

func (o *ProductProperties) GetPurchaseDateAsTime() (*time.Time, error) {
	if o.PurchaseDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.PurchaseDate, "2006-01-02T15:04:05Z07:00")
}
