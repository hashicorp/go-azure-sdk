package namespaces

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type EHNamespaceProperties struct {
	AlternateName              *string                      `json:"alternateName,omitempty"`
	ClusterArmId               *string                      `json:"clusterArmId,omitempty"`
	CreatedAt                  *string                      `json:"createdAt,omitempty"`
	DisableLocalAuth           *bool                        `json:"disableLocalAuth,omitempty"`
	Encryption                 *Encryption                  `json:"encryption,omitempty"`
	IsAutoInflateEnabled       *bool                        `json:"isAutoInflateEnabled,omitempty"`
	KafkaEnabled               *bool                        `json:"kafkaEnabled,omitempty"`
	MaximumThroughputUnits     *int64                       `json:"maximumThroughputUnits,omitempty"`
	MetricId                   *string                      `json:"metricId,omitempty"`
	PrivateEndpointConnections *[]PrivateEndpointConnection `json:"privateEndpointConnections,omitempty"`
	ProvisioningState          *string                      `json:"provisioningState,omitempty"`
	ServiceBusEndpoint         *string                      `json:"serviceBusEndpoint,omitempty"`
	Status                     *string                      `json:"status,omitempty"`
	UpdatedAt                  *string                      `json:"updatedAt,omitempty"`
	ZoneRedundant              *bool                        `json:"zoneRedundant,omitempty"`
}
