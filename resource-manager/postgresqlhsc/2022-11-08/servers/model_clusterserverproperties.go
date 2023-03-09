package servers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type ClusterServerProperties struct {
	AdministratorLogin       *string     `json:"administratorLogin,omitempty"`
	AvailabilityZone         *string     `json:"availabilityZone,omitempty"`
	CitusVersion             *string     `json:"citusVersion,omitempty"`
	EnableHa                 *bool       `json:"enableHa,omitempty"`
	EnablePublicIPAccess     *bool       `json:"enablePublicIpAccess,omitempty"`
	FullyQualifiedDomainName *string     `json:"fullyQualifiedDomainName,omitempty"`
	HaState                  *string     `json:"haState,omitempty"`
	IsReadOnly               *bool       `json:"isReadOnly,omitempty"`
	PostgresqlVersion        *string     `json:"postgresqlVersion,omitempty"`
	Role                     *ServerRole `json:"role,omitempty"`
	ServerEdition            *string     `json:"serverEdition,omitempty"`
	State                    *string     `json:"state,omitempty"`
	StorageQuotaInMb         *int64      `json:"storageQuotaInMb,omitempty"`
	VCores                   *int64      `json:"vCores,omitempty"`
}
