package namespacediscovereddevices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiscoveredMessagingEndpointsUpdate struct {
	Inbound  *map[string]DiscoveredInboundEndpointsUpdate `json:"inbound,omitempty"`
	Outbound *DiscoveredOutboundEndpointsUpdate           `json:"outbound,omitempty"`
}
