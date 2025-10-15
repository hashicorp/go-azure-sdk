package namespacediscovereddevices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiscoveredMessagingEndpoints struct {
	Inbound  *map[string]DiscoveredInboundEndpoints `json:"inbound,omitempty"`
	Outbound *DiscoveredOutboundEndpoints           `json:"outbound,omitempty"`
}
