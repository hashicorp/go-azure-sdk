package namespacedevices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessagingEndpoints struct {
	Inbound  *map[string]InboundEndpoints `json:"inbound,omitempty"`
	Outbound *OutboundEndpoints           `json:"outbound,omitempty"`
}
