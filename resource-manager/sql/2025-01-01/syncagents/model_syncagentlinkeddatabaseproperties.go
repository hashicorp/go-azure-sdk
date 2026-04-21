package syncagents

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SyncAgentLinkedDatabaseProperties struct {
	DatabaseId   *string           `json:"databaseId,omitempty"`
	DatabaseName *string           `json:"databaseName,omitempty"`
	DatabaseType *SyncMemberDbType `json:"databaseType,omitempty"`
	Description  *string           `json:"description,omitempty"`
	ServerName   *string           `json:"serverName,omitempty"`
	UserName     *string           `json:"userName,omitempty"`
}
