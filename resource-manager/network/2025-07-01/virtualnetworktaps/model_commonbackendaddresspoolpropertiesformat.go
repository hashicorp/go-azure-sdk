package virtualnetworktaps

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommonBackendAddressPoolPropertiesFormat struct {
	BackendIPConfigurations      *[]CommonNetworkInterfaceIPConfiguration    `json:"backendIPConfigurations,omitempty"`
	DrainPeriodInSeconds         *int64                                      `json:"drainPeriodInSeconds,omitempty"`
	InboundNatRules              *[]CommonSubResource                        `json:"inboundNatRules,omitempty"`
	LoadBalancerBackendAddresses *[]CommonLoadBalancerBackendAddress         `json:"loadBalancerBackendAddresses,omitempty"`
	LoadBalancingRules           *[]CommonSubResource                        `json:"loadBalancingRules,omitempty"`
	Location                     *string                                     `json:"location,omitempty"`
	OutboundRule                 *CommonSubResource                          `json:"outboundRule,omitempty"`
	OutboundRules                *[]CommonSubResource                        `json:"outboundRules,omitempty"`
	ProvisioningState            *ProvisioningState                          `json:"provisioningState,omitempty"`
	SyncMode                     *SyncMode                                   `json:"syncMode,omitempty"`
	TunnelInterfaces             *[]CommonGatewayLoadBalancerTunnelInterface `json:"tunnelInterfaces,omitempty"`
	VirtualNetwork               *CommonSubResource                          `json:"virtualNetwork,omitempty"`
}
