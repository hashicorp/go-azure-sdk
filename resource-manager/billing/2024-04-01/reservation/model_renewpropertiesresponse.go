package reservation

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RenewPropertiesResponse struct {
	BillingCurrencyTotal *Price                      `json:"billingCurrencyTotal,omitempty"`
	PricingCurrencyTotal *Price                      `json:"pricingCurrencyTotal,omitempty"`
	PurchaseProperties   *ReservationPurchaseRequest `json:"purchaseProperties,omitempty"`
}
