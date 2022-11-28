package privateendpoints

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectionProperties struct {
	GroupIds                          *[]string                              `json:"groupIds,omitempty"`
	PrivateEndpoint                   *PrivateEndpoint                       `json:"privateEndpoint"`
	PrivateLinkServiceConnectionState *ConnectionState                       `json:"privateLinkServiceConnectionState"`
	ProvisioningState                 *ConnectionPropertiesProvisioningState `json:"provisioningState,omitempty"`
}
