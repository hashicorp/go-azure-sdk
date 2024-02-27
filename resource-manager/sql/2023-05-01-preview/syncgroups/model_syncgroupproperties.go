package syncgroups

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

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

func (o *SyncGroupProperties) GetLastSyncTimeAsTime() (*time.Time, error) {
	if o.LastSyncTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastSyncTime, "2006-01-02T15:04:05Z07:00")
}
