package organizations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type LiftrBaseMarketplaceDetailsUpdate struct {
	OfferDetails       *LiftrBaseOfferDetailsUpdate   `json:"offerDetails,omitempty"`
	SubscriptionId     *string                        `json:"subscriptionId,omitempty"`
	SubscriptionStatus *MarketplaceSubscriptionStatus `json:"subscriptionStatus,omitempty"`
}
