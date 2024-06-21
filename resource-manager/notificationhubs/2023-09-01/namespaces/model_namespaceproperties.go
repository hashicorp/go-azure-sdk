package namespaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type NamespaceProperties struct {
	CreatedAt                  *string                              `json:"createdAt,omitempty"`
	Critical                   *bool                                `json:"critical,omitempty"`
	DataCenter                 *string                              `json:"dataCenter,omitempty"`
	Enabled                    *bool                                `json:"enabled,omitempty"`
	MetricId                   *string                              `json:"metricId,omitempty"`
	Name                       *string                              `json:"name,omitempty"`
	NamespaceType              *NamespaceType                       `json:"namespaceType,omitempty"`
	NetworkAcls                *NetworkAcls                         `json:"networkAcls,omitempty"`
	PnsCredentials             *PnsCredentials                      `json:"pnsCredentials,omitempty"`
	PrivateEndpointConnections *[]PrivateEndpointConnectionResource `json:"privateEndpointConnections,omitempty"`
	ProvisioningState          *OperationProvisioningState          `json:"provisioningState,omitempty"`
	PublicNetworkAccess        *PublicNetworkAccess                 `json:"publicNetworkAccess,omitempty"`
	Region                     *string                              `json:"region,omitempty"`
	ReplicationRegion          *ReplicationRegion                   `json:"replicationRegion,omitempty"`
	ScaleUnit                  *string                              `json:"scaleUnit,omitempty"`
	ServiceBusEndpoint         *string                              `json:"serviceBusEndpoint,omitempty"`
	Status                     *NamespaceStatus                     `json:"status,omitempty"`
	SubscriptionId             *string                              `json:"subscriptionId,omitempty"`
	UpdatedAt                  *string                              `json:"updatedAt,omitempty"`
	ZoneRedundancy             *ZoneRedundancyPreference            `json:"zoneRedundancy,omitempty"`
}
