package replicationfabrics

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type MasterTargetServer struct {
	AgentExpiryDate         *string            `json:"agentExpiryDate,omitempty"`
	AgentVersion            *string            `json:"agentVersion,omitempty"`
	AgentVersionDetails     *VersionDetails    `json:"agentVersionDetails,omitempty"`
	DataStores              *[]DataStore       `json:"dataStores,omitempty"`
	DiskCount               *int64             `json:"diskCount,omitempty"`
	HealthErrors            *[]HealthError     `json:"healthErrors,omitempty"`
	IPAddress               *string            `json:"ipAddress,omitempty"`
	Id                      *string            `json:"id,omitempty"`
	LastHeartbeat           *string            `json:"lastHeartbeat,omitempty"`
	MarsAgentExpiryDate     *string            `json:"marsAgentExpiryDate,omitempty"`
	MarsAgentVersion        *string            `json:"marsAgentVersion,omitempty"`
	MarsAgentVersionDetails *VersionDetails    `json:"marsAgentVersionDetails,omitempty"`
	Name                    *string            `json:"name,omitempty"`
	OsType                  *string            `json:"osType,omitempty"`
	OsVersion               *string            `json:"osVersion,omitempty"`
	RetentionVolumes        *[]RetentionVolume `json:"retentionVolumes,omitempty"`
	ValidationErrors        *[]HealthError     `json:"validationErrors,omitempty"`
	VersionStatus           *string            `json:"versionStatus,omitempty"`
}
