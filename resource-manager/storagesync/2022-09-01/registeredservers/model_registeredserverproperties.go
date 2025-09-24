package registeredservers

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type RegisteredServerProperties struct {
	ActiveAuthType             *ServerAuthType                     `json:"activeAuthType,omitempty"`
	AgentVersion               *string                             `json:"agentVersion,omitempty"`
	AgentVersionExpirationDate *string                             `json:"agentVersionExpirationDate,omitempty"`
	AgentVersionStatus         *RegisteredServerAgentVersionStatus `json:"agentVersionStatus,omitempty"`
	ApplicationId              *string                             `json:"applicationId,omitempty"`
	ClusterId                  *string                             `json:"clusterId,omitempty"`
	ClusterName                *string                             `json:"clusterName,omitempty"`
	DiscoveryEndpointUri       *string                             `json:"discoveryEndpointUri,omitempty"`
	FriendlyName               *string                             `json:"friendlyName,omitempty"`
	Identity                   *bool                               `json:"identity,omitempty"`
	LastHeartBeat              *string                             `json:"lastHeartBeat,omitempty"`
	LastOperationName          *string                             `json:"lastOperationName,omitempty"`
	LastWorkflowId             *string                             `json:"lastWorkflowId,omitempty"`
	LatestApplicationId        *string                             `json:"latestApplicationId,omitempty"`
	ManagementEndpointUri      *string                             `json:"managementEndpointUri,omitempty"`
	MonitoringConfiguration    *string                             `json:"monitoringConfiguration,omitempty"`
	MonitoringEndpointUri      *string                             `json:"monitoringEndpointUri,omitempty"`
	ProvisioningState          *string                             `json:"provisioningState,omitempty"`
	ResourceLocation           *string                             `json:"resourceLocation,omitempty"`
	ServerCertificate          *string                             `json:"serverCertificate,omitempty"`
	ServerId                   *string                             `json:"serverId,omitempty"`
	ServerManagementErrorCode  *int64                              `json:"serverManagementErrorCode,omitempty"`
	ServerName                 *string                             `json:"serverName,omitempty"`
	ServerOSVersion            *string                             `json:"serverOSVersion,omitempty"`
	ServerRole                 *string                             `json:"serverRole,omitempty"`
	ServiceLocation            *string                             `json:"serviceLocation,omitempty"`
	StorageSyncServiceUid      *string                             `json:"storageSyncServiceUid,omitempty"`
}

func (o *RegisteredServerProperties) GetAgentVersionExpirationDateAsTime() (*time.Time, error) {
	if o.AgentVersionExpirationDate == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.AgentVersionExpirationDate, "2006-01-02T15:04:05Z07:00")
}

func (o *RegisteredServerProperties) SetAgentVersionExpirationDateAsTime(input time.Time) {
	formatted := input.Format("2006-01-02T15:04:05Z07:00")
	o.AgentVersionExpirationDate = &formatted
}
