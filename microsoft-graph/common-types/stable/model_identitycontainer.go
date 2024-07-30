package stable

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityContainer struct {
	ApiConnectors      *[]IdentityApiConnector      `json:"apiConnectors,omitempty"`
	B2xUserFlows       *[]B2xIdentityUserFlow       `json:"b2xUserFlows,omitempty"`
	ConditionalAccess  *ConditionalAccessRoot       `json:"conditionalAccess,omitempty"`
	Id                 *string                      `json:"id,omitempty"`
	IdentityProviders  *[]IdentityProviderBase      `json:"identityProviders,omitempty"`
	ODataType          *string                      `json:"@odata.type,omitempty"`
	UserFlowAttributes *[]IdentityUserFlowAttribute `json:"userFlowAttributes,omitempty"`
}
