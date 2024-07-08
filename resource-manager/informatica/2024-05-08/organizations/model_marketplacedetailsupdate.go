package organizations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MarketplaceDetailsUpdate struct {
	MarketplaceSubscriptionId *string             `json:"marketplaceSubscriptionId,omitempty"`
	OfferDetails              *OfferDetailsUpdate `json:"offerDetails,omitempty"`
}
