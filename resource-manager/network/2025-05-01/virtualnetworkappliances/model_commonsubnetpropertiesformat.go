package virtualnetworkappliances

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type CommonSubnetPropertiesFormat struct {
	AddressPrefix                      *string                                          `json:"addressPrefix,omitempty"`
	AddressPrefixes                    *[]string                                        `json:"addressPrefixes,omitempty"`
	ApplicationGatewayIPConfigurations *[]CommonApplicationGatewayIPConfiguration       `json:"applicationGatewayIPConfigurations,omitempty"`
	DefaultOutboundAccess              *bool                                            `json:"defaultOutboundAccess,omitempty"`
	Delegations                        *[]CommonDelegation                              `json:"delegations,omitempty"`
	IPAllocations                      *[]CommonSubResource                             `json:"ipAllocations,omitempty"`
	IPConfigurationProfiles            *[]CommonIPConfigurationProfile                  `json:"ipConfigurationProfiles,omitempty"`
	IPConfigurations                   *[]CommonIPConfiguration                         `json:"ipConfigurations,omitempty"`
	IPamPoolPrefixAllocations          *[]CommonIPamPoolPrefixAllocation                `json:"ipamPoolPrefixAllocations,omitempty"`
	NatGateway                         *CommonSubResource                               `json:"natGateway,omitempty"`
	NetworkSecurityGroup               *CommonNetworkSecurityGroup                      `json:"networkSecurityGroup,omitempty"`
	PrivateEndpointNetworkPolicies     *VirtualNetworkPrivateEndpointNetworkPolicies    `json:"privateEndpointNetworkPolicies,omitempty"`
	PrivateEndpoints                   *[]CommonPrivateEndpoint                         `json:"privateEndpoints,omitempty"`
	PrivateLinkServiceNetworkPolicies  *VirtualNetworkPrivateLinkServiceNetworkPolicies `json:"privateLinkServiceNetworkPolicies,omitempty"`
	ProvisioningState                  *ProvisioningState                               `json:"provisioningState,omitempty"`
	Purpose                            *string                                          `json:"purpose,omitempty"`
	ResourceNavigationLinks            *[]CommonResourceNavigationLink                  `json:"resourceNavigationLinks,omitempty"`
	RouteTable                         *CommonRouteTable                                `json:"routeTable,omitempty"`
	ServiceAssociationLinks            *[]CommonServiceAssociationLink                  `json:"serviceAssociationLinks,omitempty"`
	ServiceEndpointPolicies            *[]CommonServiceEndpointPolicy                   `json:"serviceEndpointPolicies,omitempty"`
	ServiceEndpoints                   *[]CommonServiceEndpointPropertiesFormat         `json:"serviceEndpoints,omitempty"`
	ServiceGateway                     *CommonSubResource                               `json:"serviceGateway,omitempty"`
	SharingScope                       *SharingScope                                    `json:"sharingScope,omitempty"`
}
