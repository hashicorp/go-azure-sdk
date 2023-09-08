package models

// Copyright (c) HashiCorp Inc. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type IdentityContainer struct {
	ApiConnectors                    *[]IdentityApiConnector           `json:"apiConnectors,omitempty"`
	AuthenticationEventListeners     *[]AuthenticationEventListener    `json:"authenticationEventListeners,omitempty"`
	AuthenticationEventsFlows        *[]AuthenticationEventsFlow       `json:"authenticationEventsFlows,omitempty"`
	B2cUserFlows                     *[]B2cIdentityUserFlow            `json:"b2cUserFlows,omitempty"`
	B2xUserFlows                     *[]B2xIdentityUserFlow            `json:"b2xUserFlows,omitempty"`
	ConditionalAccess                *ConditionalAccessRoot            `json:"conditionalAccess,omitempty"`
	ContinuousAccessEvaluationPolicy *ContinuousAccessEvaluationPolicy `json:"continuousAccessEvaluationPolicy,omitempty"`
	CustomAuthenticationExtensions   *[]CustomAuthenticationExtension  `json:"customAuthenticationExtensions,omitempty"`
	IdentityProviders                *[]IdentityProviderBase           `json:"identityProviders,omitempty"`
	ODataType                        *string                           `json:"@odata.type,omitempty"`
	UserFlowAttributes               *[]IdentityUserFlowAttribute      `json:"userFlowAttributes,omitempty"`
	UserFlows                        *[]IdentityUserFlow               `json:"userFlows,omitempty"`
}
