package charges

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ModernChargeSummaryProperties struct {
	AzureCharges            *Amount `json:"azureCharges,omitempty"`
	BillingAccountId        *string `json:"billingAccountId,omitempty"`
	BillingPeriodId         *string `json:"billingPeriodId,omitempty"`
	BillingProfileId        *string `json:"billingProfileId,omitempty"`
	ChargesBilledSeparately *Amount `json:"chargesBilledSeparately,omitempty"`
	CustomerId              *string `json:"customerId,omitempty"`
	InvoiceSectionId        *string `json:"invoiceSectionId,omitempty"`
	IsInvoiced              *bool   `json:"isInvoiced,omitempty"`
	MarketplaceCharges      *Amount `json:"marketplaceCharges,omitempty"`
	SubscriptionId          *string `json:"subscriptionId,omitempty"`
	UsageEnd                *string `json:"usageEnd,omitempty"`
	UsageStart              *string `json:"usageStart,omitempty"`
}
