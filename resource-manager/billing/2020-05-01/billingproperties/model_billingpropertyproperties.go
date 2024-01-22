package billingproperties

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BillingPropertyProperties struct {
	AccountAdminNotificationEmailAddress *string                         `json:"accountAdminNotificationEmailAddress,omitempty"`
	BillingAccountDisplayName            *string                         `json:"billingAccountDisplayName,omitempty"`
	BillingAccountId                     *string                         `json:"billingAccountId,omitempty"`
	BillingProfileDisplayName            *string                         `json:"billingProfileDisplayName,omitempty"`
	BillingProfileId                     *string                         `json:"billingProfileId,omitempty"`
	BillingProfileSpendingLimit          *BillingProfileSpendingLimit    `json:"billingProfileSpendingLimit,omitempty"`
	BillingProfileStatus                 *BillingProfileStatus           `json:"billingProfileStatus,omitempty"`
	BillingProfileStatusReasonCode       *BillingProfileStatusReasonCode `json:"billingProfileStatusReasonCode,omitempty"`
	BillingTenantId                      *string                         `json:"billingTenantId,omitempty"`
	CostCenter                           *string                         `json:"costCenter,omitempty"`
	InvoiceSectionDisplayName            *string                         `json:"invoiceSectionDisplayName,omitempty"`
	InvoiceSectionId                     *string                         `json:"invoiceSectionId,omitempty"`
	IsAccountAdmin                       *bool                           `json:"isAccountAdmin,omitempty"`
	ProductId                            *string                         `json:"productId,omitempty"`
	ProductName                          *string                         `json:"productName,omitempty"`
	SkuDescription                       *string                         `json:"skuDescription,omitempty"`
	SkuId                                *string                         `json:"skuId,omitempty"`
}
