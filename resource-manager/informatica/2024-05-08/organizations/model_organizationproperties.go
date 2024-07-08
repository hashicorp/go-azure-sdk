package organizations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OrganizationProperties struct {
	CompanyDetails        *CompanyDetails        `json:"companyDetails,omitempty"`
	InformaticaProperties *InformaticaProperties `json:"informaticaProperties,omitempty"`
	LinkOrganization      *LinkOrganization      `json:"linkOrganization,omitempty"`
	MarketplaceDetails    *MarketplaceDetails    `json:"marketplaceDetails,omitempty"`
	ProvisioningState     *ProvisioningState     `json:"provisioningState,omitempty"`
	UserDetails           *UserDetails           `json:"userDetails,omitempty"`
}
