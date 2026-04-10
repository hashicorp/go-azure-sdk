package expressroutegateways

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoutingConfiguration struct {
	AssociatedRouteTable  *CommonSubResource    `json:"associatedRouteTable,omitempty"`
	InboundRouteMap       *CommonSubResource    `json:"inboundRouteMap,omitempty"`
	OutboundRouteMap      *CommonSubResource    `json:"outboundRouteMap,omitempty"`
	PropagatedRouteTables *PropagatedRouteTable `json:"propagatedRouteTables,omitempty"`
	VnetRoutes            *VnetRoute            `json:"vnetRoutes,omitempty"`
}
