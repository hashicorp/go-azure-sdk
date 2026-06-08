package p2svpngateways

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VnetRoute struct {
	BgpConnections     *[]CommonSubResource `json:"bgpConnections,omitempty"`
	StaticRoutes       *[]StaticRoute       `json:"staticRoutes,omitempty"`
	StaticRoutesConfig *StaticRoutesConfig  `json:"staticRoutesConfig,omitempty"`
}
