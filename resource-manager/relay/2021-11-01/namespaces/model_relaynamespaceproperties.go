package namespaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RelayNamespaceProperties struct {
	CreatedAt                  *string                      `json:"createdAt,omitempty"`
	MetricId                   *string                      `json:"metricId,omitempty"`
	PrivateEndpointConnections *[]PrivateEndpointConnection `json:"privateEndpointConnections,omitempty"`
	ProvisioningState          *string                      `json:"provisioningState,omitempty"`
	PublicNetworkAccess        *PublicNetworkAccess         `json:"publicNetworkAccess,omitempty"`
	ServiceBusEndpoint         *string                      `json:"serviceBusEndpoint,omitempty"`
	Status                     *string                      `json:"status,omitempty"`
	UpdatedAt                  *string                      `json:"updatedAt,omitempty"`
}
