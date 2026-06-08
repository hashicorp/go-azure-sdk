package applicationgateways

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommonNetworkInterfaceIPConfigurationPropertiesFormat struct {
	ApplicationGatewayBackendAddressPools *[]CommonApplicationGatewayBackendAddressPool                         `json:"applicationGatewayBackendAddressPools,omitempty"`
	ApplicationSecurityGroups             *[]CommonApplicationSecurityGroup                                     `json:"applicationSecurityGroups,omitempty"`
	GatewayLoadBalancer                   *CommonSubResource                                                    `json:"gatewayLoadBalancer,omitempty"`
	LoadBalancerBackendAddressPools       *[]CommonBackendAddressPool                                           `json:"loadBalancerBackendAddressPools,omitempty"`
	LoadBalancerInboundNatRules           *[]CommonInboundNatRule                                               `json:"loadBalancerInboundNatRules,omitempty"`
	Primary                               *bool                                                                 `json:"primary,omitempty"`
	PrivateIPAddress                      *string                                                               `json:"privateIPAddress,omitempty"`
	PrivateIPAddressPrefixLength          *int64                                                                `json:"privateIPAddressPrefixLength,omitempty"`
	PrivateIPAddressVersion               *IPVersion                                                            `json:"privateIPAddressVersion,omitempty"`
	PrivateIPAllocationMethod             *IPAllocationMethod                                                   `json:"privateIPAllocationMethod,omitempty"`
	PrivateLinkConnectionProperties       *CommonNetworkInterfaceIPConfigurationPrivateLinkConnectionProperties `json:"privateLinkConnectionProperties,omitempty"`
	ProvisioningState                     *ProvisioningState                                                    `json:"provisioningState,omitempty"`
	PublicIPAddress                       *CommonPublicIPAddress                                                `json:"publicIPAddress,omitempty"`
	Subnet                                *CommonSubnet                                                         `json:"subnet,omitempty"`
	VirtualNetworkTaps                    *[]CommonVirtualNetworkTap                                            `json:"virtualNetworkTaps,omitempty"`
}
