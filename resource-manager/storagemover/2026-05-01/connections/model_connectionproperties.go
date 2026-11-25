package connections

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ConnectionProperties struct {
	ConnectionStatus          *ConnectionStatus  `json:"connectionStatus,omitempty"`
	Description               *string            `json:"description,omitempty"`
	JobList                   *[]string          `json:"jobList,omitempty"`
	PrivateEndpointName       *string            `json:"privateEndpointName,omitempty"`
	PrivateEndpointResourceId *string            `json:"privateEndpointResourceId,omitempty"`
	PrivateLinkServiceId      string             `json:"privateLinkServiceId"`
	ProvisioningState         *ProvisioningState `json:"provisioningState,omitempty"`
}
