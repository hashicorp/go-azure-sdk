package namespacediscovereddevices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NamespaceDiscoveredDeviceProperties struct {
	Attributes             *map[string]interface{}       `json:"attributes,omitempty"`
	DiscoveryId            string                        `json:"discoveryId"`
	Endpoints              *DiscoveredMessagingEndpoints `json:"endpoints,omitempty"`
	ExternalDeviceId       *string                       `json:"externalDeviceId,omitempty"`
	Manufacturer           *string                       `json:"manufacturer,omitempty"`
	Model                  *string                       `json:"model,omitempty"`
	OperatingSystem        *string                       `json:"operatingSystem,omitempty"`
	OperatingSystemVersion *string                       `json:"operatingSystemVersion,omitempty"`
	ProvisioningState      *ProvisioningState            `json:"provisioningState,omitempty"`
	Version                int64                         `json:"version"`
}
