package privateendpointconnections

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RemotePrivateEndpointConnection struct {
	PrivateEndpoint                   *ArmIdWrapper               `json:"privateEndpoint,omitempty"`
	PrivateLinkServiceConnectionState *PrivateLinkConnectionState `json:"privateLinkServiceConnectionState,omitempty"`
	ProvisioningState                 *string                     `json:"provisioningState,omitempty"`
}
