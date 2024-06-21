package replicas

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerProperties struct {
	AdministratorLogin       *string                `json:"administratorLogin,omitempty"`
	EarliestRestoreDate      *string                `json:"earliestRestoreDate,omitempty"`
	FullyQualifiedDomainName *string                `json:"fullyQualifiedDomainName,omitempty"`
	MasterServerId           *string                `json:"masterServerId,omitempty"`
	MinimalTlsVersion        *MinimalTlsVersionEnum `json:"minimalTlsVersion,omitempty"`
	ReplicaCapacity          *int64                 `json:"replicaCapacity,omitempty"`
	ReplicationRole          *string                `json:"replicationRole,omitempty"`
	SslEnforcement           *SslEnforcementEnum    `json:"sslEnforcement,omitempty"`
	StorageProfile           *StorageProfile        `json:"storageProfile,omitempty"`
	UserVisibleState         *ServerState           `json:"userVisibleState,omitempty"`
	Version                  *ServerVersion         `json:"version,omitempty"`
}
