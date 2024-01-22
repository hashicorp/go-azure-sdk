package billingsubscriptions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingSubscriptionProperties struct {
	BillingProfileDisplayName *string                        `json:"billingProfileDisplayName,omitempty"`
	BillingProfileId          *string                        `json:"billingProfileId,omitempty"`
	CustomerDisplayName       *string                        `json:"customerDisplayName,omitempty"`
	CustomerId                *string                        `json:"customerId,omitempty"`
	DisplayName               *string                        `json:"displayName,omitempty"`
	InvoiceSectionDisplayName *string                        `json:"invoiceSectionDisplayName,omitempty"`
	InvoiceSectionId          *string                        `json:"invoiceSectionId,omitempty"`
	LastMonthCharges          *Amount                        `json:"lastMonthCharges,omitempty"`
	MonthToDateCharges        *Amount                        `json:"monthToDateCharges,omitempty"`
	Reseller                  *Reseller                      `json:"reseller,omitempty"`
	SkuDescription            *string                        `json:"skuDescription,omitempty"`
	SkuId                     *string                        `json:"skuId,omitempty"`
	SubscriptionBillingStatus *BillingSubscriptionStatusType `json:"subscriptionBillingStatus,omitempty"`
	SubscriptionId            *string                        `json:"subscriptionId,omitempty"`
}
