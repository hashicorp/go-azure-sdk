package webapplicationfirewallpolicies

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ApplicationGatewayPrivateEndpointConnectionProperties struct {
	LinkIdentifier                    *string                                  `json:"linkIdentifier,omitempty"`
	PrivateEndpoint                   *CommonPrivateEndpoint                   `json:"privateEndpoint,omitempty"`
	PrivateLinkServiceConnectionState *CommonPrivateLinkServiceConnectionState `json:"privateLinkServiceConnectionState,omitempty"`
	ProvisioningState                 *ProvisioningState                       `json:"provisioningState,omitempty"`
}
