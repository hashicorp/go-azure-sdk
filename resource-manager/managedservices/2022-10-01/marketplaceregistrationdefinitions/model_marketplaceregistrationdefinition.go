package marketplaceregistrationdefinitions

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MarketplaceRegistrationDefinition struct {
	Id         *string                                      `json:"id,omitempty"`
	Name       *string                                      `json:"name,omitempty"`
	Plan       *Plan                                        `json:"plan"`
	Properties *MarketplaceRegistrationDefinitionProperties `json:"properties"`
	Type       *string                                      `json:"type,omitempty"`
}
