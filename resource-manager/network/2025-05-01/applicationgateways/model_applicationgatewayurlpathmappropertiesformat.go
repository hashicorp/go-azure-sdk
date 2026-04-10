package applicationgateways

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationGatewayURLPathMapPropertiesFormat struct {
	DefaultBackendAddressPool     *CommonSubResource            `json:"defaultBackendAddressPool,omitempty"`
	DefaultBackendHTTPSettings    *CommonSubResource            `json:"defaultBackendHttpSettings,omitempty"`
	DefaultLoadDistributionPolicy *CommonSubResource            `json:"defaultLoadDistributionPolicy,omitempty"`
	DefaultRedirectConfiguration  *CommonSubResource            `json:"defaultRedirectConfiguration,omitempty"`
	DefaultRewriteRuleSet         *CommonSubResource            `json:"defaultRewriteRuleSet,omitempty"`
	PathRules                     *[]ApplicationGatewayPathRule `json:"pathRules,omitempty"`
	ProvisioningState             *ProvisioningState            `json:"provisioningState,omitempty"`
}
