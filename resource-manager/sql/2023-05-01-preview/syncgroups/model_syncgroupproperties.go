package syncgroups

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SyncGroupProperties struct {
	ConflictLoggingRetentionInDays *int64                        `json:"conflictLoggingRetentionInDays,omitempty"`
	ConflictResolutionPolicy       *SyncConflictResolutionPolicy `json:"conflictResolutionPolicy,omitempty"`
	EnableConflictLogging          *bool                         `json:"enableConflictLogging,omitempty"`
	HubDatabasePassword            *string                       `json:"hubDatabasePassword,omitempty"`
	HubDatabaseUserName            *string                       `json:"hubDatabaseUserName,omitempty"`
	Interval                       *int64                        `json:"interval,omitempty"`
	LastSyncTime                   *string                       `json:"lastSyncTime,omitempty"`
	PrivateEndpointName            *string                       `json:"privateEndpointName,omitempty"`
	Schema                         *SyncGroupSchema              `json:"schema,omitempty"`
	SyncDatabaseId                 *string                       `json:"syncDatabaseId,omitempty"`
	SyncState                      *SyncGroupState               `json:"syncState,omitempty"`
	UsePrivateLinkConnection       *bool                         `json:"usePrivateLinkConnection,omitempty"`
}
