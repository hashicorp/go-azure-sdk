package networkclouds

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NetworkConfiguration struct {
	AttachedNetworkConfiguration        *AttachedNetworkConfiguration        `json:"attachedNetworkConfiguration,omitempty"`
	BgpServiceLoadBalancerConfiguration *BgpServiceLoadBalancerConfiguration `json:"bgpServiceLoadBalancerConfiguration,omitempty"`
	CloudServicesNetworkId              string                               `json:"cloudServicesNetworkId"`
	CniNetworkId                        string                               `json:"cniNetworkId"`
	DnsServiceIP                        *string                              `json:"dnsServiceIp,omitempty"`
	L2ServiceLoadBalancerConfiguration  *L2ServiceLoadBalancerConfiguration  `json:"l2ServiceLoadBalancerConfiguration,omitempty"`
	PodCidrs                            *[]string                            `json:"podCidrs,omitempty"`
	ServiceCidrs                        *[]string                            `json:"serviceCidrs,omitempty"`
}
