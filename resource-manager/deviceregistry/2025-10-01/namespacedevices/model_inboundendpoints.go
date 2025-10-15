package namespacedevices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type InboundEndpoints struct {
	AdditionalConfiguration *string             `json:"additionalConfiguration,omitempty"`
	Address                 string              `json:"address"`
	Authentication          *HostAuthentication `json:"authentication,omitempty"`
	EndpointType            string              `json:"endpointType"`
	TrustSettings           *TrustSettings      `json:"trustSettings,omitempty"`
	Version                 *string             `json:"version,omitempty"`
}
