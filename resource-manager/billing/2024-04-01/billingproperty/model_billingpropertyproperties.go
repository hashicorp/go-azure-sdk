package billingproperty

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingPropertyProperties struct {
	AccountAdminNotificationEmailAddress *string                             `json:"accountAdminNotificationEmailAddress,omitempty"`
	BillingAccountAgreementType          *AgreementType                      `json:"billingAccountAgreementType,omitempty"`
	BillingAccountDisplayName            *string                             `json:"billingAccountDisplayName,omitempty"`
	BillingAccountId                     *string                             `json:"billingAccountId,omitempty"`
	BillingAccountSoldToCountry          *string                             `json:"billingAccountSoldToCountry,omitempty"`
	BillingAccountStatus                 *AccountStatus                      `json:"billingAccountStatus,omitempty"`
	BillingAccountStatusReasonCode       *BillingAccountStatusReasonCode     `json:"billingAccountStatusReasonCode,omitempty"`
	BillingAccountSubType                *AccountSubType                     `json:"billingAccountSubType,omitempty"`
	BillingAccountType                   *AccountType                        `json:"billingAccountType,omitempty"`
	BillingCurrency                      *string                             `json:"billingCurrency,omitempty"`
	BillingProfileDisplayName            *string                             `json:"billingProfileDisplayName,omitempty"`
	BillingProfileId                     *string                             `json:"billingProfileId,omitempty"`
	BillingProfilePaymentMethodFamily    *PaymentMethodFamily                `json:"billingProfilePaymentMethodFamily,omitempty"`
	BillingProfilePaymentMethodType      *string                             `json:"billingProfilePaymentMethodType,omitempty"`
	BillingProfileSpendingLimit          *SpendingLimit                      `json:"billingProfileSpendingLimit,omitempty"`
	BillingProfileSpendingLimitDetails   *[]SpendingLimitDetails             `json:"billingProfileSpendingLimitDetails,omitempty"`
	BillingProfileStatus                 *BillingProfileStatus               `json:"billingProfileStatus,omitempty"`
	BillingProfileStatusReasonCode       *BillingProfileStatusReasonCode     `json:"billingProfileStatusReasonCode,omitempty"`
	BillingTenantId                      *string                             `json:"billingTenantId,omitempty"`
	CostCenter                           *string                             `json:"costCenter,omitempty"`
	CustomerDisplayName                  *string                             `json:"customerDisplayName,omitempty"`
	CustomerId                           *string                             `json:"customerId,omitempty"`
	CustomerStatus                       *CustomerStatus                     `json:"customerStatus,omitempty"`
	EnrollmentDetails                    *SubscriptionEnrollmentDetails      `json:"enrollmentDetails,omitempty"`
	InvoiceSectionDisplayName            *string                             `json:"invoiceSectionDisplayName,omitempty"`
	InvoiceSectionId                     *string                             `json:"invoiceSectionId,omitempty"`
	InvoiceSectionStatus                 *InvoiceSectionState                `json:"invoiceSectionStatus,omitempty"`
	InvoiceSectionStatusReasonCode       *InvoiceSectionStateReasonCode      `json:"invoiceSectionStatusReasonCode,omitempty"`
	IsAccountAdmin                       *bool                               `json:"isAccountAdmin,omitempty"`
	IsTransitionedBillingAccount         *bool                               `json:"isTransitionedBillingAccount,omitempty"`
	ProductId                            *string                             `json:"productId,omitempty"`
	ProductName                          *string                             `json:"productName,omitempty"`
	SkuDescription                       *string                             `json:"skuDescription,omitempty"`
	SkuId                                *string                             `json:"skuId,omitempty"`
	SubscriptionBillingStatus            *BillingSubscriptionStatus          `json:"subscriptionBillingStatus,omitempty"`
	SubscriptionBillingStatusDetails     *[]BillingSubscriptionStatusDetails `json:"subscriptionBillingStatusDetails,omitempty"`
	SubscriptionBillingType              *SubscriptionBillingType            `json:"subscriptionBillingType,omitempty"`
	SubscriptionServiceUsageAddress      *AddressDetails                     `json:"subscriptionServiceUsageAddress,omitempty"`
	SubscriptionWorkloadType             *SubscriptionWorkloadType           `json:"subscriptionWorkloadType,omitempty"`
}
