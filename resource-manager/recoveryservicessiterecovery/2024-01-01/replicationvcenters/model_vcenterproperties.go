package replicationvcenters

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type VCenterProperties struct {
	DiscoveryStatus       *string        `json:"discoveryStatus,omitempty"`
	FabricArmResourceName *string        `json:"fabricArmResourceName,omitempty"`
	FriendlyName          *string        `json:"friendlyName,omitempty"`
	HealthErrors          *[]HealthError `json:"healthErrors,omitempty"`
	IPAddress             *string        `json:"ipAddress,omitempty"`
	InfrastructureId      *string        `json:"infrastructureId,omitempty"`
	InternalId            *string        `json:"internalId,omitempty"`
	LastHeartbeat         *string        `json:"lastHeartbeat,omitempty"`
	Port                  *string        `json:"port,omitempty"`
	ProcessServerId       *string        `json:"processServerId,omitempty"`
	RunAsAccountId        *string        `json:"runAsAccountId,omitempty"`
}
