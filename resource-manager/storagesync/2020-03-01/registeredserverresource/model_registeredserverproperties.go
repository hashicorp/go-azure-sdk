package registeredserverresource

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegisteredServerProperties struct {
	AgentVersion               *string                             `json:"agentVersion,omitempty"`
	AgentVersionExpirationDate *string                             `json:"agentVersionExpirationDate,omitempty"`
	AgentVersionStatus         *RegisteredServerAgentVersionStatus `json:"agentVersionStatus,omitempty"`
	ClusterId                  *string                             `json:"clusterId,omitempty"`
	ClusterName                *string                             `json:"clusterName,omitempty"`
	DiscoveryEndpointUri       *string                             `json:"discoveryEndpointUri,omitempty"`
	FriendlyName               *string                             `json:"friendlyName,omitempty"`
	LastHeartBeat              *string                             `json:"lastHeartBeat,omitempty"`
	LastOperationName          *string                             `json:"lastOperationName,omitempty"`
	LastWorkflowId             *string                             `json:"lastWorkflowId,omitempty"`
	ManagementEndpointUri      *string                             `json:"managementEndpointUri,omitempty"`
	MonitoringConfiguration    *string                             `json:"monitoringConfiguration,omitempty"`
	MonitoringEndpointUri      *string                             `json:"monitoringEndpointUri,omitempty"`
	ProvisioningState          *string                             `json:"provisioningState,omitempty"`
	ResourceLocation           *string                             `json:"resourceLocation,omitempty"`
	ServerCertificate          *string                             `json:"serverCertificate,omitempty"`
	ServerId                   *string                             `json:"serverId,omitempty"`
	ServerManagementErrorCode  *int64                              `json:"serverManagementErrorCode,omitempty"`
	ServerOSVersion            *string                             `json:"serverOSVersion,omitempty"`
	ServerRole                 *string                             `json:"serverRole,omitempty"`
	ServiceLocation            *string                             `json:"serviceLocation,omitempty"`
	StorageSyncServiceUid      *string                             `json:"storageSyncServiceUid,omitempty"`
}
