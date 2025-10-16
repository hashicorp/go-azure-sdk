package namespacediscovereddevices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NamespaceDiscoveredDeviceUpdateProperties struct {
	Attributes             *map[string]interface{}             `json:"attributes,omitempty"`
	DiscoveryId            *string                             `json:"discoveryId,omitempty"`
	Endpoints              *DiscoveredMessagingEndpointsUpdate `json:"endpoints,omitempty"`
	ExternalDeviceId       *string                             `json:"externalDeviceId,omitempty"`
	OperatingSystemVersion *string                             `json:"operatingSystemVersion,omitempty"`
	Version                *int64                              `json:"version,omitempty"`
}
