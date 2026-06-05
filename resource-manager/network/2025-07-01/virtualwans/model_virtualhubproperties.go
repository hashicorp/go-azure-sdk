package virtualwans

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VirtualHubProperties struct {
	AddressPrefix                       *string                              `json:"addressPrefix,omitempty"`
	AllowBranchToBranchTraffic          *bool                                `json:"allowBranchToBranchTraffic,omitempty"`
	AzureFirewall                       *CommonSubResource                   `json:"azureFirewall,omitempty"`
	BgpConnections                      *[]CommonSubResource                 `json:"bgpConnections,omitempty"`
	ExpressRouteGateway                 *CommonSubResource                   `json:"expressRouteGateway,omitempty"`
	HubRoutingPreference                *HubRoutingPreference                `json:"hubRoutingPreference,omitempty"`
	IPConfigurations                    *[]CommonSubResource                 `json:"ipConfigurations,omitempty"`
	P2SVpnGateway                       *CommonSubResource                   `json:"p2SVpnGateway,omitempty"`
	PreferredRoutingGateway             *PreferredRoutingGateway             `json:"preferredRoutingGateway,omitempty"`
	ProvisioningState                   *ProvisioningState                   `json:"provisioningState,omitempty"`
	RouteMaps                           *[]CommonSubResource                 `json:"routeMaps,omitempty"`
	RouteTable                          *VirtualHubRouteTable                `json:"routeTable,omitempty"`
	RoutingState                        *RoutingState                        `json:"routingState,omitempty"`
	SecurityPartnerProvider             *CommonSubResource                   `json:"securityPartnerProvider,omitempty"`
	SecurityProviderName                *string                              `json:"securityProviderName,omitempty"`
	Sku                                 *string                              `json:"sku,omitempty"`
	VirtualHubRouteTableV2s             *[]VirtualHubRouteTableV2            `json:"virtualHubRouteTableV2s,omitempty"`
	VirtualRouterAsn                    *int64                               `json:"virtualRouterAsn,omitempty"`
	VirtualRouterAutoScaleConfiguration *VirtualRouterAutoScaleConfiguration `json:"virtualRouterAutoScaleConfiguration,omitempty"`
	VirtualRouterIPs                    *[]string                            `json:"virtualRouterIps,omitempty"`
	VirtualWAN                          *CommonSubResource                   `json:"virtualWan,omitempty"`
	VpnGateway                          *CommonSubResource                   `json:"vpnGateway,omitempty"`
}
