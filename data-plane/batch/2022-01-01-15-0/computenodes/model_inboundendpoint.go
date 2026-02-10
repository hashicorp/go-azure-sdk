package computenodes

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InboundEndpoint struct {
	BackendPort     int64                   `json:"backendPort"`
	FrontendPort    int64                   `json:"frontendPort"`
	Name            string                  `json:"name"`
	Protocol        InboundEndpointProtocol `json:"protocol"`
	PublicFQDN      string                  `json:"publicFQDN"`
	PublicIPAddress string                  `json:"publicIPAddress"`
}
