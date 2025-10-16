package namespacedevices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type OutboundEndpointsUpdate struct {
	Assigned   *map[string]DeviceMessagingEndpointUpdate `json:"assigned,omitempty"`
	Unassigned *map[string]DeviceMessagingEndpointUpdate `json:"unassigned,omitempty"`
}
