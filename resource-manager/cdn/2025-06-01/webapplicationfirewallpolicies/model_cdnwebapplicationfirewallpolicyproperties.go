package webapplicationfirewallpolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CdnWebApplicationFirewallPolicyProperties struct {
	CustomRules        *CustomRuleList      `json:"customRules,omitempty"`
	EndpointLinks      *[]CdnEndpoint       `json:"endpointLinks,omitempty"`
	ExtendedProperties *map[string]string   `json:"extendedProperties,omitempty"`
	ManagedRules       *ManagedRuleSetList  `json:"managedRules,omitempty"`
	PolicySettings     *PolicySettings      `json:"policySettings,omitempty"`
	ProvisioningState  *ProvisioningState   `json:"provisioningState,omitempty"`
	RateLimitRules     *RateLimitRuleList   `json:"rateLimitRules,omitempty"`
	ResourceState      *PolicyResourceState `json:"resourceState,omitempty"`
}
