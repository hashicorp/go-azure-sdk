package vaults

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type PrivateEndpointConnection struct {
	GroupIds                          *[]VaultSubResourceType            `json:"groupIds,omitempty"`
	PrivateEndpoint                   *PrivateEndpoint                   `json:"privateEndpoint"`
	PrivateLinkServiceConnectionState *PrivateLinkServiceConnectionState `json:"privateLinkServiceConnectionState"`
	ProvisioningState                 *ProvisioningState                 `json:"provisioningState,omitempty"`
}
