package organizations

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OrganizationPropertiesCustomUpdate struct {
	CompanyDetails                    *CompanyDetailsUpdate                  `json:"companyDetails,omitempty"`
	ExistingResourceId                *string                                `json:"existingResourceId,omitempty"`
	InformaticaOrganizationProperties *InformaticaOrganizationResourceUpdate `json:"informaticaOrganizationProperties,omitempty"`
	MarketplaceDetails                *MarketplaceDetailsUpdate              `json:"marketplaceDetails,omitempty"`
	UserDetails                       *UserDetailsUpdate                     `json:"userDetails,omitempty"`
}
