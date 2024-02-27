package syncagents

import (
	"time"

	"github.com/hashicorp/go-azure-helpers/lang/dates"
)

// Copyright (c) Microsoft Corporation. All rights reserved.
// Licensed under the MIT License. See NOTICE.txt in the project root for license information.

type SyncAgentProperties struct {
	ExpiryTime     *string         `json:"expiryTime,omitempty"`
	IsUpToDate     *bool           `json:"isUpToDate,omitempty"`
	LastAliveTime  *string         `json:"lastAliveTime,omitempty"`
	Name           *string         `json:"name,omitempty"`
	State          *SyncAgentState `json:"state,omitempty"`
	SyncDatabaseId *string         `json:"syncDatabaseId,omitempty"`
	Version        *string         `json:"version,omitempty"`
}

func (o *SyncAgentProperties) GetExpiryTimeAsTime() (*time.Time, error) {
	if o.ExpiryTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.ExpiryTime, "2006-01-02T15:04:05Z07:00")
}

func (o *SyncAgentProperties) GetLastAliveTimeAsTime() (*time.Time, error) {
	if o.LastAliveTime == nil {
		return nil, nil
	}
	return dates.ParseAsFormat(o.LastAliveTime, "2006-01-02T15:04:05Z07:00")
}
