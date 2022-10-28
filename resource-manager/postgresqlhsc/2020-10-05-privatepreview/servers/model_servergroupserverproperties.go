package servers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ServerGroupServerProperties struct {
	AdministratorLogin       *string            `json:"administratorLogin,omitempty"`
	AvailabilityZone         *Zone              `json:"availabilityZone,omitempty"`
	CitusVersion             *CitusVersion      `json:"citusVersion,omitempty"`
	EnableHa                 *bool              `json:"enableHa,omitempty"`
	EnablePublicIP           *bool              `json:"enablePublicIp,omitempty"`
	FullyQualifiedDomainName *string            `json:"fullyQualifiedDomainName,omitempty"`
	HaState                  *ServerHaState     `json:"haState,omitempty"`
	PostgresqlVersion        *PostgreSQLVersion `json:"postgresqlVersion,omitempty"`
	Role                     *ServerRole        `json:"role,omitempty"`
	ServerEdition            *ServerEdition     `json:"serverEdition,omitempty"`
	StandbyAvailabilityZone  *string            `json:"standbyAvailabilityZone,omitempty"`
	State                    *ServerState       `json:"state,omitempty"`
	StorageQuotaInMb         *int64             `json:"storageQuotaInMb,omitempty"`
	VCores                   *int64             `json:"vCores,omitempty"`
}
