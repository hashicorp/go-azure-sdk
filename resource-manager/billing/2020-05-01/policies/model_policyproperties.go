package policies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PolicyProperties struct {
	MarketplacePurchases *MarketplacePurchasesPolicy `json:"marketplacePurchases,omitempty"`
	ReservationPurchases *ReservationPurchasesPolicy `json:"reservationPurchases,omitempty"`
	ViewCharges          *ViewChargesPolicy          `json:"viewCharges,omitempty"`
}
