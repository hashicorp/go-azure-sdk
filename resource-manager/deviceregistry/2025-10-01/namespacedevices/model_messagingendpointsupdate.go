package namespacedevices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessagingEndpointsUpdate struct {
	Inbound  *map[string]InboundEndpointsUpdate `json:"inbound,omitempty"`
	Outbound *OutboundEndpointsUpdate           `json:"outbound,omitempty"`
}
