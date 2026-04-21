package syncmembers

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SyncMemberProperties struct {
	DatabaseName                      *string           `json:"databaseName,omitempty"`
	DatabaseType                      *SyncMemberDbType `json:"databaseType,omitempty"`
	Password                          *string           `json:"password,omitempty"`
	PrivateEndpointName               *string           `json:"privateEndpointName,omitempty"`
	ServerName                        *string           `json:"serverName,omitempty"`
	SqlServerDatabaseId               *string           `json:"sqlServerDatabaseId,omitempty"`
	SyncAgentId                       *string           `json:"syncAgentId,omitempty"`
	SyncDirection                     *SyncDirection    `json:"syncDirection,omitempty"`
	SyncMemberAzureDatabaseResourceId *string           `json:"syncMemberAzureDatabaseResourceId,omitempty"`
	SyncState                         *SyncMemberState  `json:"syncState,omitempty"`
	UsePrivateLinkConnection          *bool             `json:"usePrivateLinkConnection,omitempty"`
	UserName                          *string           `json:"userName,omitempty"`
}
