package loadbalancers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommonLoadBalancerPropertiesFormat struct {
	BackendAddressPools      *[]CommonBackendAddressPool      `json:"backendAddressPools,omitempty"`
	FrontendIPConfigurations *[]CommonFrontendIPConfiguration `json:"frontendIPConfigurations,omitempty"`
	InboundNatPools          *[]CommonInboundNatPool          `json:"inboundNatPools,omitempty"`
	InboundNatRules          *[]CommonInboundNatRule          `json:"inboundNatRules,omitempty"`
	LoadBalancingRules       *[]CommonLoadBalancingRule       `json:"loadBalancingRules,omitempty"`
	OutboundRules            *[]CommonOutboundRule            `json:"outboundRules,omitempty"`
	Probes                   *[]CommonProbe                   `json:"probes,omitempty"`
	ProvisioningState        *ProvisioningState               `json:"provisioningState,omitempty"`
	ResourceGuid             *string                          `json:"resourceGuid,omitempty"`
	Scope                    *LoadBalancerScope               `json:"scope,omitempty"`
}
