package billingsubscriptionsaliases

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingSubscriptionAliasProperties struct {
	AutoRenew                            *AutoRenew                            `json:"autoRenew,omitempty"`
	BeneficiaryTenantId                  *string                               `json:"beneficiaryTenantId,omitempty"`
	BillingFrequency                     *string                               `json:"billingFrequency,omitempty"`
	BillingPolicies                      *map[string]string                    `json:"billingPolicies,omitempty"`
	BillingProfileDisplayName            *string                               `json:"billingProfileDisplayName,omitempty"`
	BillingProfileId                     *string                               `json:"billingProfileId,omitempty"`
	BillingProfileName                   *string                               `json:"billingProfileName,omitempty"`
	BillingSubscriptionId                *string                               `json:"billingSubscriptionId,omitempty"`
	ConsumptionCostCenter                *string                               `json:"consumptionCostCenter,omitempty"`
	CustomerDisplayName                  *string                               `json:"customerDisplayName,omitempty"`
	CustomerId                           *string                               `json:"customerId,omitempty"`
	CustomerName                         *string                               `json:"customerName,omitempty"`
	DisplayName                          *string                               `json:"displayName,omitempty"`
	EnrollmentAccountDisplayName         *string                               `json:"enrollmentAccountDisplayName,omitempty"`
	EnrollmentAccountId                  *string                               `json:"enrollmentAccountId,omitempty"`
	EnrollmentAccountSubscriptionDetails *EnrollmentAccountSubscriptionDetails `json:"enrollmentAccountSubscriptionDetails,omitempty"`
	InvoiceSectionDisplayName            *string                               `json:"invoiceSectionDisplayName,omitempty"`
	InvoiceSectionId                     *string                               `json:"invoiceSectionId,omitempty"`
	InvoiceSectionName                   *string                               `json:"invoiceSectionName,omitempty"`
	LastMonthCharges                     *Amount                               `json:"lastMonthCharges,omitempty"`
	MonthToDateCharges                   *Amount                               `json:"monthToDateCharges,omitempty"`
	NextBillingCycleDetails              *NextBillingCycleDetails              `json:"nextBillingCycleDetails,omitempty"`
	OfferId                              *string                               `json:"offerId,omitempty"`
	ProductCategory                      *string                               `json:"productCategory,omitempty"`
	ProductType                          *string                               `json:"productType,omitempty"`
	ProductTypeId                        *string                               `json:"productTypeId,omitempty"`
	PurchaseDate                         *string                               `json:"purchaseDate,omitempty"`
	Quantity                             *int64                                `json:"quantity,omitempty"`
	RenewalTermDetails                   *RenewalTermDetails                   `json:"renewalTermDetails,omitempty"`
	Reseller                             *Reseller                             `json:"reseller,omitempty"`
	SkuDescription                       *string                               `json:"skuDescription,omitempty"`
	SkuId                                *string                               `json:"skuId,omitempty"`
	Status                               *BillingSubscriptionStatus            `json:"status,omitempty"`
	SubscriptionId                       *string                               `json:"subscriptionId,omitempty"`
	SuspensionReasons                    *[]string                             `json:"suspensionReasons,omitempty"`
	TermDuration                         *string                               `json:"termDuration,omitempty"`
	TermEndDate                          *string                               `json:"termEndDate,omitempty"`
	TermStartDate                        *string                               `json:"termStartDate,omitempty"`
}
