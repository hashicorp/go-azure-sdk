package marketplaceregistrationdefinitions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MarketplaceRegistrationDefinitionProperties struct {
	Authorizations         []Authorization          `json:"authorizations"`
	EligibleAuthorizations *[]EligibleAuthorization `json:"eligibleAuthorizations,omitempty"`
	ManagedByTenantId      string                   `json:"managedByTenantId"`
	OfferDisplayName       *string                  `json:"offerDisplayName,omitempty"`
	PlanDisplayName        *string                  `json:"planDisplayName,omitempty"`
	PublisherDisplayName   *string                  `json:"publisherDisplayName,omitempty"`
}
