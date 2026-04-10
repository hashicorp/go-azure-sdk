package networksecuritygroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommonFrontendIPConfigurationPropertiesFormat struct {
	GatewayLoadBalancer       *CommonSubResource     `json:"gatewayLoadBalancer,omitempty"`
	InboundNatPools           *[]CommonSubResource   `json:"inboundNatPools,omitempty"`
	InboundNatRules           *[]CommonSubResource   `json:"inboundNatRules,omitempty"`
	LoadBalancingRules        *[]CommonSubResource   `json:"loadBalancingRules,omitempty"`
	OutboundRules             *[]CommonSubResource   `json:"outboundRules,omitempty"`
	PrivateIPAddress          *string                `json:"privateIPAddress,omitempty"`
	PrivateIPAddressVersion   *IPVersion             `json:"privateIPAddressVersion,omitempty"`
	PrivateIPAllocationMethod *IPAllocationMethod    `json:"privateIPAllocationMethod,omitempty"`
	ProvisioningState         *ProvisioningState     `json:"provisioningState,omitempty"`
	PublicIPAddress           *CommonPublicIPAddress `json:"publicIPAddress,omitempty"`
	PublicIPPrefix            *CommonSubResource     `json:"publicIPPrefix,omitempty"`
	Subnet                    *CommonSubnet          `json:"subnet,omitempty"`
}
