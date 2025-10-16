package akriconnectortemplate

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type AkriConnectorTemplateDeviceInboundEndpointType struct {
	DisplayName  *string `json:"displayName,omitempty"`
	EndpointType string  `json:"endpointType"`
	Version      *string `json:"version,omitempty"`
}
