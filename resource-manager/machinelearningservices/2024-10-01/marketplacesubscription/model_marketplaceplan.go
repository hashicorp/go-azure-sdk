package marketplacesubscription

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MarketplacePlan struct {
	OfferId     *string `json:"offerId,omitempty"`
	PlanId      *string `json:"planId,omitempty"`
	PublisherId *string `json:"publisherId,omitempty"`
}
