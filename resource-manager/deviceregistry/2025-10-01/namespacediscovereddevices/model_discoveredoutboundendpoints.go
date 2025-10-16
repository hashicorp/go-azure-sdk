package namespacediscovereddevices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DiscoveredOutboundEndpoints struct {
	Assigned map[string]DeviceMessagingEndpoint `json:"assigned"`
}
