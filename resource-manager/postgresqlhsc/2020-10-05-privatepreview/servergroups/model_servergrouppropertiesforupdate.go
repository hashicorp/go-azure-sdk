package servergroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerGroupPropertiesForUpdate struct {
	AdministratorLoginPassword *string            `json:"administratorLoginPassword,omitempty"`
	AvailabilityZone           *string            `json:"availabilityZone,omitempty"`
	BackupRetentionDays        *int64             `json:"backupRetentionDays,omitempty"`
	CitusVersion               *CitusVersion      `json:"citusVersion,omitempty"`
	EnableShardsOnCoordinator  *bool              `json:"enableShardsOnCoordinator,omitempty"`
	MaintenanceWindow          *MaintenanceWindow `json:"maintenanceWindow"`
	PostgresqlVersion          *PostgreSQLVersion `json:"postgresqlVersion,omitempty"`
	ServerRoleGroups           *[]ServerRoleGroup `json:"serverRoleGroups,omitempty"`
	StandbyAvailabilityZone    *string            `json:"standbyAvailabilityZone,omitempty"`
}
