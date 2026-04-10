package loadbalancers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommonLoadBalancerBackendAddressPropertiesFormat struct {
	AdminState                          *LoadBalancerBackendAddressAdminState `json:"adminState,omitempty"`
	IPAddress                           *string                               `json:"ipAddress,omitempty"`
	InboundNatRulesPortMapping          *[]CommonNatRulePortMapping           `json:"inboundNatRulesPortMapping,omitempty"`
	LoadBalancerFrontendIPConfiguration *CommonSubResource                    `json:"loadBalancerFrontendIPConfiguration,omitempty"`
	NetworkInterfaceIPConfiguration     *CommonSubResource                    `json:"networkInterfaceIPConfiguration,omitempty"`
	Subnet                              *CommonSubResource                    `json:"subnet,omitempty"`
	VirtualNetwork                      *CommonSubResource                    `json:"virtualNetwork,omitempty"`
}
