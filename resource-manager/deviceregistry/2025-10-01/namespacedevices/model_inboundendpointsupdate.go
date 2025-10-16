package namespacedevices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InboundEndpointsUpdate struct {
	AdditionalConfiguration *string                   `json:"additionalConfiguration,omitempty"`
	Address                 *string                   `json:"address,omitempty"`
	Authentication          *HostAuthenticationUpdate `json:"authentication,omitempty"`
	EndpointType            *string                   `json:"endpointType,omitempty"`
	TrustSettings           *TrustSettings            `json:"trustSettings,omitempty"`
	Version                 *string                   `json:"version,omitempty"`
}
