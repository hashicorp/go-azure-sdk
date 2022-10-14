package charges

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LegacyChargeSummaryProperties struct {
	AzureCharges            *float64 `json:"azureCharges,omitempty"`
	AzureMarketplaceCharges *float64 `json:"azureMarketplaceCharges,omitempty"`
	BillingPeriodId         *string  `json:"billingPeriodId,omitempty"`
	ChargesBilledSeparately *float64 `json:"chargesBilledSeparately,omitempty"`
	Currency                *string  `json:"currency,omitempty"`
	UsageEnd                *string  `json:"usageEnd,omitempty"`
	UsageStart              *string  `json:"usageStart,omitempty"`
}
