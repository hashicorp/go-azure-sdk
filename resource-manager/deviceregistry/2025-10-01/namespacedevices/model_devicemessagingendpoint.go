package namespacedevices

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type DeviceMessagingEndpoint struct {
	Address      string  `json:"address"`
	EndpointType *string `json:"endpointType,omitempty"`
}
