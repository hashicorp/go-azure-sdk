package marketplacesubscription

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MarketplaceSubscription struct {
	MarketplacePlan               *MarketplacePlan                          `json:"marketplacePlan,omitempty"`
	MarketplaceSubscriptionStatus *MarketplaceSubscriptionStatus            `json:"marketplaceSubscriptionStatus,omitempty"`
	ModelId                       string                                    `json:"modelId"`
	ProvisioningState             *MarketplaceSubscriptionProvisioningState `json:"provisioningState,omitempty"`
}
