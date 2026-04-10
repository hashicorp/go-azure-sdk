package webapplicationfirewallpolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationGatewayPathRulePropertiesFormat struct {
	BackendAddressPool     *CommonSubResource `json:"backendAddressPool,omitempty"`
	BackendHTTPSettings    *CommonSubResource `json:"backendHttpSettings,omitempty"`
	FirewallPolicy         *CommonSubResource `json:"firewallPolicy,omitempty"`
	LoadDistributionPolicy *CommonSubResource `json:"loadDistributionPolicy,omitempty"`
	Paths                  *[]string          `json:"paths,omitempty"`
	ProvisioningState      *ProvisioningState `json:"provisioningState,omitempty"`
	RedirectConfiguration  *CommonSubResource `json:"redirectConfiguration,omitempty"`
	RewriteRuleSet         *CommonSubResource `json:"rewriteRuleSet,omitempty"`
}
