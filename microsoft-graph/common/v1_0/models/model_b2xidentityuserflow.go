package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type B2xIdentityUserFlow struct {
	ApiConnectorConfiguration *UserFlowApiConnectorConfiguration     `json:"apiConnectorConfiguration,omitempty"`
	Id                        *string                                `json:"id,omitempty"`
	IdentityProviders         *[]IdentityProvider                    `json:"identityProviders,omitempty"`
	Languages                 *[]UserFlowLanguageConfiguration       `json:"languages,omitempty"`
	ODataType                 *string                                `json:"@odata.type,omitempty"`
	UserAttributeAssignments  *[]IdentityUserFlowAttributeAssignment `json:"userAttributeAssignments,omitempty"`
	UserFlowIdentityProviders *[]IdentityProviderBase                `json:"userFlowIdentityProviders,omitempty"`
	UserFlowType              *B2xIdentityUserFlowUserFlowType       `json:"userFlowType,omitempty"`
}
