package beta

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type B2cIdentityUserFlow struct {
	ApiConnectorConfiguration      *UserFlowApiConnectorConfiguration     `json:"apiConnectorConfiguration,omitempty"`
	DefaultLanguageTag             *string                                `json:"defaultLanguageTag,omitempty"`
	Id                             *string                                `json:"id,omitempty"`
	IdentityProviders              *[]IdentityProvider                    `json:"identityProviders,omitempty"`
	IsLanguageCustomizationEnabled *bool                                  `json:"isLanguageCustomizationEnabled,omitempty"`
	Languages                      *[]UserFlowLanguageConfiguration       `json:"languages,omitempty"`
	ODataType                      *string                                `json:"@odata.type,omitempty"`
	UserAttributeAssignments       *[]IdentityUserFlowAttributeAssignment `json:"userAttributeAssignments,omitempty"`
	UserFlowIdentityProviders      *[]IdentityProviderBase                `json:"userFlowIdentityProviders,omitempty"`
	UserFlowType                   *B2cIdentityUserFlowUserFlowType       `json:"userFlowType,omitempty"`
	UserFlowTypeVersion            *float64                               `json:"userFlowTypeVersion,omitempty"`
}
