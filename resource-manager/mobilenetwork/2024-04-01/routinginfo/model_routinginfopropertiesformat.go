package routinginfo

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RoutingInfoPropertiesFormat struct {
	ControlPlaneAccessRoutes *[]IPv4Route                                     `json:"controlPlaneAccessRoutes,omitempty"`
	UserPlaneAccessRoutes    *[]IPv4Route                                     `json:"userPlaneAccessRoutes,omitempty"`
	UserPlaneDataRoutes      *[]UserPlaneDataRoutesUserPlaneDataRoutesInlined `json:"userPlaneDataRoutes,omitempty"`
}
