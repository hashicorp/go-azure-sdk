package servicegateways

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommonVirtualNetworkPropertiesFormat struct {
	AddressSpace                *CommonAddressSpace                 `json:"addressSpace,omitempty"`
	BgpCommunities              *CommonVirtualNetworkBgpCommunities `json:"bgpCommunities,omitempty"`
	DdosProtectionPlan          *CommonSubResource                  `json:"ddosProtectionPlan,omitempty"`
	DefaultPublicNatGateway     *CommonSubResource                  `json:"defaultPublicNatGateway,omitempty"`
	DhcpOptions                 *CommonDhcpOptions                  `json:"dhcpOptions,omitempty"`
	EnableDdosProtection        *bool                               `json:"enableDdosProtection,omitempty"`
	EnableVMProtection          *bool                               `json:"enableVmProtection,omitempty"`
	Encryption                  *CommonVirtualNetworkEncryption     `json:"encryption,omitempty"`
	FlowLogs                    *[]CommonFlowLog                    `json:"flowLogs,omitempty"`
	FlowTimeoutInMinutes        *int64                              `json:"flowTimeoutInMinutes,omitempty"`
	IPAllocations               *[]CommonSubResource                `json:"ipAllocations,omitempty"`
	PrivateEndpointVNetPolicies *PrivateEndpointVNetPolicies        `json:"privateEndpointVNetPolicies,omitempty"`
	ProvisioningState           *ProvisioningState                  `json:"provisioningState,omitempty"`
	ResourceGuid                *string                             `json:"resourceGuid,omitempty"`
	Subnets                     *[]CommonSubnet                     `json:"subnets,omitempty"`
	VirtualNetworkPeerings      *[]CommonVirtualNetworkPeering      `json:"virtualNetworkPeerings,omitempty"`
}
