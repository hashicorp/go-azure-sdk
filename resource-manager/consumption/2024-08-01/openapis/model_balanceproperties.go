package openapis

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type BalanceProperties struct {
	AdjustmentDetails              *[]BalancePropertiesAdjustmentDetailsItem   `json:"adjustmentDetails,omitempty"`
	Adjustments                    *float64                                    `json:"adjustments,omitempty"`
	AzureMarketplaceServiceCharges *float64                                    `json:"azureMarketplaceServiceCharges,omitempty"`
	BeginningBalance               *float64                                    `json:"beginningBalance,omitempty"`
	BillingFrequency               *BillingFrequency                           `json:"billingFrequency,omitempty"`
	ChargesBilledSeparately        *float64                                    `json:"chargesBilledSeparately,omitempty"`
	Currency                       *string                                     `json:"currency,omitempty"`
	EndingBalance                  *float64                                    `json:"endingBalance,omitempty"`
	NewPurchases                   *float64                                    `json:"newPurchases,omitempty"`
	NewPurchasesDetails            *[]BalancePropertiesNewPurchasesDetailsItem `json:"newPurchasesDetails,omitempty"`
	OverageRefund                  *float64                                    `json:"overageRefund,omitempty"`
	PriceHidden                    *bool                                       `json:"priceHidden,omitempty"`
	ServiceOverage                 *float64                                    `json:"serviceOverage,omitempty"`
	TotalOverage                   *float64                                    `json:"totalOverage,omitempty"`
	TotalUsage                     *float64                                    `json:"totalUsage,omitempty"`
	Utilized                       *float64                                    `json:"utilized,omitempty"`
}
