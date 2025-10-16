package namespaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MessagingEndpoint struct {
	Address      string  `json:"address"`
	EndpointType *string `json:"endpointType,omitempty"`
	ResourceId   *string `json:"resourceId,omitempty"`
}
