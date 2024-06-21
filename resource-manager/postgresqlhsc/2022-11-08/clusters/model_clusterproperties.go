package clusters

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterProperties struct {
	AdministratorLogin              *string                            `json:"administratorLogin,omitempty"`
	AdministratorLoginPassword      *string                            `json:"administratorLoginPassword,omitempty"`
	CitusVersion                    *string                            `json:"citusVersion,omitempty"`
	CoordinatorEnablePublicIPAccess *bool                              `json:"coordinatorEnablePublicIpAccess,omitempty"`
	CoordinatorServerEdition        *string                            `json:"coordinatorServerEdition,omitempty"`
	CoordinatorStorageQuotaInMb     *int64                             `json:"coordinatorStorageQuotaInMb,omitempty"`
	CoordinatorVCores               *int64                             `json:"coordinatorVCores,omitempty"`
	EarliestRestoreTime             *string                            `json:"earliestRestoreTime,omitempty"`
	EnableHa                        *bool                              `json:"enableHa,omitempty"`
	EnableShardsOnCoordinator       *bool                              `json:"enableShardsOnCoordinator,omitempty"`
	MaintenanceWindow               *MaintenanceWindow                 `json:"maintenanceWindow,omitempty"`
	NodeCount                       *int64                             `json:"nodeCount,omitempty"`
	NodeEnablePublicIPAccess        *bool                              `json:"nodeEnablePublicIpAccess,omitempty"`
	NodeServerEdition               *string                            `json:"nodeServerEdition,omitempty"`
	NodeStorageQuotaInMb            *int64                             `json:"nodeStorageQuotaInMb,omitempty"`
	NodeVCores                      *int64                             `json:"nodeVCores,omitempty"`
	PointInTimeUTC                  *string                            `json:"pointInTimeUTC,omitempty"`
	PostgresqlVersion               *string                            `json:"postgresqlVersion,omitempty"`
	PreferredPrimaryZone            *string                            `json:"preferredPrimaryZone,omitempty"`
	PrivateEndpointConnections      *[]SimplePrivateEndpointConnection `json:"privateEndpointConnections,omitempty"`
	ProvisioningState               *string                            `json:"provisioningState,omitempty"`
	ReadReplicas                    *[]string                          `json:"readReplicas,omitempty"`
	ServerNames                     *[]ServerNameItem                  `json:"serverNames,omitempty"`
	SourceLocation                  *string                            `json:"sourceLocation,omitempty"`
	SourceResourceId                *string                            `json:"sourceResourceId,omitempty"`
	State                           *string                            `json:"state,omitempty"`
}
